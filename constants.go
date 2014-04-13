package libvirt

/*
#cgo LDFLAGS: -lvirt -ldl
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

// virDomainState
const (
	VIR_DOMAIN_NOSTATE     = C.VIR_DOMAIN_NOSTATE
	VIR_DOMAIN_RUNNING     = C.VIR_DOMAIN_RUNNING
	VIR_DOMAIN_BLOCKED     = C.VIR_DOMAIN_BLOCKED
	VIR_DOMAIN_PAUSED      = C.VIR_DOMAIN_PAUSED
	VIR_DOMAIN_SHUTDOWN    = C.VIR_DOMAIN_SHUTDOWN
	VIR_DOMAIN_CRASHED     = C.VIR_DOMAIN_CRASHED
	VIR_DOMAIN_PMSUSPENDED = C.VIR_DOMAIN_PMSUSPENDED
	VIR_DOMAIN_SHUTOFF     = C.VIR_DOMAIN_SHUTOFF
)

// virDomainMetadataType
const (
	VIR_DOMAIN_METADATA_DESCRIPTION = C.VIR_DOMAIN_METADATA_DESCRIPTION
	VIR_DOMAIN_METADATA_TITLE       = C.VIR_DOMAIN_METADATA_TITLE
	VIR_DOMAIN_METADATA_ELEMENT     = C.VIR_DOMAIN_METADATA_ELEMENT
)
