# go-pinproc

Interface to the libpinproc library from within Go.

This library is used communicate with a [Pinball - Remote Operations Controller](https://www.multimorphic.com/store/circuit-boards/p-roc/) (P-ROC) from Multimorphic, Inc.

It provides an interface to the C library found here:

https://github.com/preble/libpinproc

Use at your own risk. This code is being used with a Judge Dredd pinball
machine and it is available here for those who might find it useful. It has
not been tested thoroughly, it is under active development, and it is not
guaranteed to work on any other machine. Review this code before use. It is
your responsibility to ensure the safety of your own pinball machine.

The `examples` directory contains code to test the various capabilities of the P-ROC. To blink all the lamps in the pinball machine:

```
go run examples/lamp-test/lamp-test.go
```

## License

MIT