package out

import (
	"github.com/miclip/dotnet-resource"
)

type Request struct {
	Source dotnetresource.Source `json:"source"`
	Params Params            `json:"params"`
}

type Params struct {
	Project        string `json:"project"`
}

type Response struct {
	Version  dotnetresource.Version        `json:"version"`
	Metadata []dotnetresource.MetadataPair `json:"metadata"`
}