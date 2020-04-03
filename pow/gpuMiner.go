package pow

//this command generates a go file containing the complete opencl source as a string constant
//allows the opencl sources to be burned into the miner executable
//AND fixes the problem with nvidia's compute cache not respecting #include statements
//go:generate go run generate_opencl.go

import (
	"fmt"
	"github.com/charliehorse55/go-opencl/cl"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/util"
	"math/big"
	"time"
	"unsafe"
)

//the constant needs have a width of hash + divisor
//src: https://arxiv.org/pdf/1902.01961.pdf
//hash is sha256 - 32 bytes. divisor varies. round it up to the nearest 4 byte boundary
//smaller values save time checking the result on the gpu
//right now its 2 which makes the max difficulty 2^64-1
//if this is exceeded, just increase it to 3!
const (
	numDivisorWords = 2
	divisorConstantWords = 8 + numDivisorWords
	divisorConstantBytes = divisorConstantWords * 4
)

var gpuLog = util.NewLogger("pow", "gpuMiner")

type GpuMiner struct {
	config.GPUConfig

	//opencl driver state
	context *cl.Context
	queue *cl.CommandQueue
	kernel *cl.Kernel

	name string

	//nvidia's openCL driver forces you to use spinlocks
	//solution: sleep for a while before calling back into openCL's spinlock
	sleepTime time.Duration

	//gpu buffers
	prefix, mulDivisor, output *cl.MemObject
}

func GetOpenCLGPUs() ([]*cl.Device, error) {
	platforms, err := cl.GetPlatforms()
	if err != nil {
		if err.Error() == "cl: error -1001" {
			gpuLog.Info("No OpenCL platforms avilable")
			return nil, nil
		}
		return nil, err
	}
	gpus := []*cl.Device{}
	for _,platform := range platforms {
		devices, err := platform.GetDevices(cl.DeviceTypeGPU)
		if err != nil {
			if err.Error() == "cl: Device Not Found" {
				gpuLog.Info("Found 0 GPUs on platform %s\n", platform.Name())
				continue
			}
			return nil, fmt.Errorf("failed to get devices for platform %s: %s", platform.Name(), err.Error())
		}
		gpus = append(gpus, devices...)
	}
	return gpus, nil
}

//burn a #define into the top of the kernel
//if you pass this in a as kernel parameter, the compiler can't unroll loops as effectively
func GenKernelSource() string {
	return fmt.Sprintf("#define MUL_CONSTANT_SIZE (%d)\n%s", divisorConstantWords, KernelSource)
}

//prepare an openCL device for work
func NewGpuMiner(device *cl.Device, config *config.GPUConfig, poolEnabled bool) (*GpuMiner, error) {

	var g GpuMiner
	var err error
	if config == nil {
		g.Count = 8
		g.GroupSize = 64
		g.Groups = 4096
	} else {
		g.GPUConfig = *config
	}
	g.name = device.Name()
	g.context, err = cl.CreateContext([]*cl.Device{device})
	if err != nil {
		return nil, fmt.Errorf("CreateContext failed: %+v", err)
	}
	g.queue, err = g.context.CreateCommandQueue(device, 0)
	if err != nil {
		return nil, fmt.Errorf("CreateCommandQueue failed: %+v", err)
	}
	program, err := g.context.CreateProgramWithSource([]string{GenKernelSource()})
	if err != nil {
		return nil, fmt.Errorf("CreateProgramWithSource failed: %+v", err)
	}
	if err := program.BuildProgram(nil, "-Werror"); err != nil {
		return nil, fmt.Errorf("BuildProgram failed: %+v", err)
	}
	g.kernel, err = program.CreateKernel("tellor")
	if err != nil {
		return nil, fmt.Errorf("CreateKernel failed: %+v", err)
	}
	if poolEnabled {
		g.prefix, err = g.context.CreateEmptyBuffer(cl.MemReadOnly, 64)
	} else {
		g.prefix, err = g.context.CreateEmptyBuffer(cl.MemReadOnly, 56)
	}
	if err != nil {
		return nil, fmt.Errorf("CreateBuffer failed for prefix: %+v", err)
	}
	g.mulDivisor, err = g.context.CreateEmptyBuffer(cl.MemReadOnly, 128)
	if err != nil {
		return nil, fmt.Errorf("CreateBuffer failed for mulDivisor: %+v", err)
	}
	g.output, err = g.context.CreateEmptyBuffer(cl.MemWriteOnly, 16)
	if err != nil {
		return nil, fmt.Errorf("CreateBuffer failed for output: %+v", err)
	}
	return &g, nil
}

func (g *GpuMiner)Name() string {
	return g.name
}

func (g *GpuMiner)CheckRange(hash *HashSettings,  start uint64, n uint64) (string, uint64, error) {
	if n % g.StepSize() != 0 {
		return "", 0, fmt.Errorf("n (%d) must be a multiple of GPU step size (%d)", n, g.StepSize())
	}
	mulDivisorBytes := createDivisorByteArray(hash.difficulty)

	_, err := g.queue.EnqueueWriteBuffer(g.prefix, true, 0, len(hash.prefix), unsafe.Pointer(&hash.prefix[0]), nil)
	if err != nil {
		return "", 0, fmt.Errorf("EnqueueWriteBuffer hashPrefix failed: %+v", err)
	}
	_, err = g.queue.EnqueueWriteBuffer(g.mulDivisor, true, 0, len(mulDivisorBytes), unsafe.Pointer(&mulDivisorBytes[0]), nil)
	if err != nil {
		return "", 0, fmt.Errorf("EnqueueWriteBuffer mulDivisor failed: %+v", err)
	}

	done := uint64(0)
	for done < n {
		if err := g.kernel.SetArgs(g.prefix, g.mulDivisor, g.output, start, g.Count); err != nil {
			return "", done, fmt.Errorf("SetKernelArgs failed: %+v", err)
		}

		kernelStarted := time.Now()
		_, err := g.queue.EnqueueNDRangeKernel(g.kernel, nil, []int{g.Groups*g.GroupSize}, []int{g.GroupSize}, nil)
		if err != nil {
			return "", done, fmt.Errorf("EnqueueNDRangeKernel failed: %+v", err)
		}
		//flush the q then sleep while we wait for the kernel to finish
		g.queue.Flush()
		time.Sleep(g.sleepTime)

		//kernel will be done soon, start waiting for it to finish with ReadBuffer
		//on nvidia hardware this spinlocks, on AMD hardware it sleeps
		readStarted := time.Now()
		//16 byte nonce generated by gpu
		results := make([]byte, 16)
		_, err = g.queue.EnqueueReadBuffer(g.output, true, 0, len(results), unsafe.Pointer(&results[0]), nil)
		if err != nil {
			return "", done, fmt.Errorf("EnqueueReadBuffer failed: %+v", err)
		}
		end := time.Now()
		totalTime := end.Sub(kernelStarted)
		readTime := end.Sub(readStarted)
		readTarget := (totalTime * 10)/100
		g.sleepTime += readTime - readTarget

		start += g.StepSize()
		done += g.StepSize()
		if results[0] != 0 {
			return string(results), done, nil
		}
	}
	return "", done, nil
}

//number of hashes this backend checks at a time
func (g *GpuMiner)StepSize() uint64 {
	return uint64(g.Groups)*uint64(g.GroupSize)*uint64(g.Count)
}

func fullBigInt(n int) *big.Int {
	b := make([]byte, n, n)
	for i := 0; i < n; i++ {
		b[i] = 0xff
	}
	x := new(big.Int)
	x.SetBytes(b)
	return x
}

func littleEndianPad(x *big.Int, n int) []byte {
	b := x.Bytes()
	if len(b) < n {
		c := make([]byte, n-len(b))
		b = append(c, b...)
	}
	for i := 0; i < len(b)/2; i++ {
		tmp := b[i]
		b[i] = b[len(b)-(i+1)]
		b[len(b)-(i+1)] = tmp
	}
	return b
}

// quickly check divisibility using a multiply by a precomputed constant
// see https://lemire.me/blog/2019/02/08/faster-remainders-when-the-divisor-is-a-constant-beating-compilers-and-libdivide/
func createDivisorByteArray(d *big.Int) []byte {

	length := divisorConstantBytes

	c := fullBigInt(length)
	bigone := big.NewInt(1)
	c.Div(c, d)
	c.Add(c, bigone)
	cSub1 := new(big.Int)
	cSub1.Sub(c, bigone)

	result := make([]byte, 2*length, 2*length)
	copy(result, littleEndianPad(c, length))
	copy(result[length:], littleEndianPad(cSub1, length))

	return result
}
