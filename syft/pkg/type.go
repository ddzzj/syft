package pkg

import (
	"github.com/anchore/packageurl-go"
)

// Type represents a Package Type for or within a language ecosystem (there may be multiple package types within a language ecosystem)
type Type string

const (
	// the full set of supported packages
	UnknownPkg            Type = "UnknownPackage"
	AlpmPkg               Type = "alpm"
	ApkPkg                Type = "apk"
	BinaryPkg             Type = "binary"
	CocoapodsPkg          Type = "pod"
	ConanPkg              Type = "conan"
	DartPubPkg            Type = "dart-pub"
	DebPkg                Type = "deb"
	DotnetPkg             Type = "dotnet"
	GemPkg                Type = "gem"
	GoModulePkg           Type = "go-module"
	GraalVMNativeImagePkg Type = "graalvm-native-image"
	HackagePkg            Type = "hackage"
	HexPkg                Type = "hex"
	JavaPkg               Type = "java-archive"
	JenkinsPluginPkg      Type = "jenkins-plugin"
	KbPkg                 Type = "msrc-kb"
	NixPkg                Type = "nix"
	NpmPkg                Type = "npm"
	PhpComposerPkg        Type = "php-composer"
	PortagePkg            Type = "portage"
	PythonPkg             Type = "python"
	RpmPkg                Type = "rpm"
	RustPkg               Type = "rust-crate"
)

// AllPkgs represents all supported package types
var AllPkgs = []Type{
	AlpmPkg,
	ApkPkg,
	BinaryPkg,
	CocoapodsPkg,
	ConanPkg,
	DartPubPkg,
	DebPkg,
	DotnetPkg,
	GemPkg,
	GoModulePkg,
	HackagePkg,
	HexPkg,
	JavaPkg,
	JenkinsPluginPkg,
	KbPkg,
	NixPkg,
	NpmPkg,
	PhpComposerPkg,
	PortagePkg,
	PythonPkg,
	RpmPkg,
	RustPkg,
}

// PackageURLType returns the PURL package type for the current package.
func (t Type) PackageURLType() string {
	switch t {
	case AlpmPkg:
		return "alpm"
	case ApkPkg:
		return packageurl.TypeAlpine
	case CocoapodsPkg:
		return packageurl.TypeCocoapods
	case ConanPkg:
		return packageurl.TypeConan
	case DartPubPkg:
		return packageurl.TypePub
	case DebPkg:
		return "deb"
	case DotnetPkg:
		return packageurl.TypeDotnet
	case GemPkg:
		return packageurl.TypeGem
	case HexPkg:
		return packageurl.TypeHex
	case GoModulePkg:
		return packageurl.TypeGolang
	case HackagePkg:
		return packageurl.TypeHackage
	case JavaPkg, JenkinsPluginPkg:
		return packageurl.TypeMaven
	case PhpComposerPkg:
		return packageurl.TypeComposer
	case PythonPkg:
		return packageurl.TypePyPi
	case PortagePkg:
		return "portage"
	case NixPkg:
		return "nix"
	case NpmPkg:
		return packageurl.TypeNPM
	case RpmPkg:
		return packageurl.TypeRPM
	case RustPkg:
		return "cargo"
	default:
		// TODO: should this be a "generic" purl type instead?
		return ""
	}
}

func TypeFromPURL(p string) Type {
	purl, err := packageurl.FromString(p)
	if err != nil {
		return UnknownPkg
	}

	return TypeByName(purl.Type)
}

func TypeByName(name string) Type {
	switch name {
	case packageurl.TypeDebian:
		return DebPkg
	case packageurl.TypeRPM:
		return RpmPkg
	case "alpm":
		return AlpmPkg
	case packageurl.TypeAlpine, "alpine":
		return ApkPkg
	case packageurl.TypeMaven:
		return JavaPkg
	case packageurl.TypeComposer:
		return PhpComposerPkg
	case packageurl.TypeGolang:
		return GoModulePkg
	case packageurl.TypeNPM:
		return NpmPkg
	case packageurl.TypePyPi:
		return PythonPkg
	case packageurl.TypeGem:
		return GemPkg
	case "cargo", "crate":
		return RustPkg
	case packageurl.TypePub:
		return DartPubPkg
	case packageurl.TypeDotnet:
		return DotnetPkg
	case packageurl.TypeCocoapods:
		return CocoapodsPkg
	case packageurl.TypeConan:
		return ConanPkg
	case packageurl.TypeHackage:
		return HackagePkg
	case "portage":
		return PortagePkg
	case packageurl.TypeHex:
		return HexPkg
	case "nix":
		return NixPkg
	default:
		return UnknownPkg
	}
}
