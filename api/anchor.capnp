using Go = import "/go.capnp";

@0xefb5a91f96d44de3;

$Go.package("anchor");
$Go.import("github.com/wetware/ww/internal/api/anchor");


# NOTE:  the Value API is unstable and may change withouth
#        warning.  Do not use in production settings.
struct Value {
    union {
        nil                @0 :Void;
        capability         @1 :Capability;
        chan                  :union {
            sender         @2 :import "channel.capnp".Sender;
            recver         @3 :import "channel.capnp".PeekRecver;
            closer         @4 :import "channel.capnp".Closer;
            sendCloser     @5 :import "channel.capnp".SendCloser;
            sendRecvCloser @6 :import "channel.capnp".PeekableChan;
        }

        # TODO(soon):  process, string, []byte, numerical, AnyPointer, etc.
    }
}

interface Loader {
    load  @0 () -> (value :Value);
}

interface Storer {
    store @0 (value :Value) -> ();
}

interface Register extends(Loader, Storer) {}


interface Anchor {
    ls   @0 () -> (children :List(Child));
    walk @1 (path :Text) -> (anchor :Anchor);

    struct Child {
        name   @0 :Text;
        anchor @1 :Anchor;
    }
}
