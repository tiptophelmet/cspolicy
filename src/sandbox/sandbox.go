package sandbox

type SandboxVal string

const (
	AllowDownloads                      SandboxVal = "allow-downloads"
	AllowForms                          SandboxVal = "allow-forms"
	AllowModals                         SandboxVal = "allow-modals"
	AllowOrientationLock                SandboxVal = "allow-orientation-lock"
	AllowPointerLock                    SandboxVal = "allow-pointer-lock"
	AllowPopups                         SandboxVal = "allow-popups"
	AllowPopupsToEscapeSandbox          SandboxVal = "allow-popups-to-escape-sandbox"
	AllowPresentation                   SandboxVal = "allow-presentation"
	AllowSameOrigin                     SandboxVal = "allow-same-origin"
	AllowScripts                        SandboxVal = "allow-scripts"
	AllowTopNavigation                  SandboxVal = "allow-top-navigation"
	AllowTopNavigationByUserActivation  SandboxVal = "allow-top-navigation-by-user-activation"
	AllowTopNavigationToCustomProtocols SandboxVal = "allow-top-navigation-to-custom-protocols"
)
