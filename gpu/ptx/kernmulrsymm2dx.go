package ptx

//This file is auto-generated. Editing is futile.

func init() { Code["kernmulrsymm2dx"] = KERNMULRSYMM2DX }

const KERNMULRSYMM2DX = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_000038cc_00000000-9_kernmulrsymm2dx.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/nimble-cube/gpu/ptx/kernmulrsymm2dx.cu"

.visible .entry kernmulRSymm2Dx(
	.param .u64 kernmulRSymm2Dx_param_0,
	.param .u64 kernmulRSymm2Dx_param_1,
	.param .u32 kernmulRSymm2Dx_param_2,
	.param .u32 kernmulRSymm2Dx_param_3
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<26>;
	.reg .f32 	%f<6>;
	.reg .s64 	%rd<11>;


	ld.param.u64 	%rd3, [kernmulRSymm2Dx_param_0];
	ld.param.u64 	%rd4, [kernmulRSymm2Dx_param_1];
	ld.param.u32 	%r3, [kernmulRSymm2Dx_param_2];
	ld.param.u32 	%r4, [kernmulRSymm2Dx_param_3];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 19 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	.loc 2 20 1
	mov.u32 	%r8, %ntid.x;
	mov.u32 	%r9, %ctaid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	.loc 2 22 1
	setp.ge.s32 	%p1, %r2, %r4;
	setp.ge.s32 	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	.loc 2 26 1
	mad.lo.s32 	%r11, %r1, %r4, %r2;
	.loc 2 27 1
	sub.s32 	%r12, %r3, %r1;
	mad.lo.s32 	%r13, %r12, %r4, %r2;
	.loc 2 31 1
	shr.u32 	%r14, %r3, 31;
	add.s32 	%r15, %r3, %r14;
	shr.s32 	%r16, %r15, 1;
	add.s32 	%r17, %r16, 1;
	setp.lt.s32 	%p4, %r1, %r17;
	.loc 2 32 1
	selp.b32 	%r18, %r11, %r13, %p4;
	mul.wide.s32 	%rd5, %r18, 4;
	add.s64 	%rd6, %rd2, %rd5;
	.loc 2 37 1
	shl.b32 	%r19, %r11, 1;
	.loc 2 39 1
	mul.wide.s32 	%rd7, %r19, 4;
	add.s64 	%rd8, %rd1, %rd7;
	.loc 2 40 1
	add.s32 	%r20, %r19, 1;
	mul.wide.s32 	%rd9, %r20, 4;
	add.s64 	%rd10, %rd1, %rd9;
	ld.global.f32 	%f1, [%rd10];
	.loc 2 39 1
	ld.global.f32 	%f2, [%rd8];
	.loc 2 37 1
	ld.global.f32 	%f3, [%rd6];
	.loc 2 42 1
	mul.f32 	%f4, %f2, %f3;
	st.global.f32 	[%rd8], %f4;
	.loc 2 43 1
	mul.f32 	%f5, %f1, %f3;
	st.global.f32 	[%rd10], %f5;

BB0_2:
	.loc 2 44 2
	ret;
}


`
