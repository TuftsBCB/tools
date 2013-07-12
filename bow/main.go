package main

import (
	"github.com/TuftsBCB/frags/bow"
	"github.com/TuftsBCB/tools/util"
)

func init() {
	util.FlagUse("cpu")
	util.FlagParse("frag-lib-dir chain pdb-file out-bow", "")
	util.AssertNArg(4)
}

func main() {
	libPath := util.Arg(0)
	chain := util.Arg(1)
	pdbEntryPath := util.Arg(2)
	bowOut := util.Arg(3)

	lib := util.StructureLibrary(libPath)
	entry := util.PDBRead(pdbEntryPath)

	thechain := entry.Chain(chain[0])
	if thechain == nil || !thechain.IsProtein() {
		util.Fatalf("Could not find chain with identifier '%c'.", chain[0])
	}

	bow := bow.StructureBOW(lib, thechain)
	util.BOWWrite(util.CreateFile(bowOut), bow)
}
