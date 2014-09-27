namespace-utils
===============

Utilities for Linux Kernel Namespaces

  nsview : This will show the IDs of various namespaces running on a host.

===============

Build :  (Ensure that you have Go/Golang installed)

    $go build nsview.go

  

===============

Run : #./nsview

    IPC namespaces [4026531839 4026532455 4026532566 4026532677]

    MNT namespaces [4026531840 4026531856 4026532307 4026532325]

    NET namespaces [4026531956 4026532221 4026532458 4026532569]

    PID namespaces [4026531836 4026532456 4026532567 4026532678]

    UTS namespaces [4026531838 4026532454 4026532565 4026532676]

===============
