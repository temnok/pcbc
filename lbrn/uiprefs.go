// Copyright © 2025 Alex Temnok. All rights reserved.

package lbrn

type UIPrefs struct {
	Optimize_ByLayer           *Param
	Optimize_ByGroup           *Param
	Optimize_ByPriority        *Param
	Optimize_WhichDirection    *Param
	Optimize_InnerToOuter      *Param
	Optimize_ByDirection       *Param
	Optimize_ReduceTravel      *Param
	Optimize_HideBacklash      *Param
	Optimize_ReduceDirChanges  *Param
	Optimize_ChooseCorners     *Param
	Optimize_AllowReverse      *Param
	Optimize_RemoveOverlaps    *Param
	Optimize_OptimalEntryPoint *Param
	Optimize_OverlapDist       *Param
}

var UIPrefsDefaults = &UIPrefs{
	Optimize_ByLayer:           &Param{Value: "0"},
	Optimize_ByGroup:           &Param{Value: "-1"},
	Optimize_ByPriority:        &Param{Value: "1"},
	Optimize_WhichDirection:    &Param{Value: "0"},
	Optimize_InnerToOuter:      &Param{Value: "1"},
	Optimize_ByDirection:       &Param{Value: "0"},
	Optimize_ReduceTravel:      &Param{Value: "1"},
	Optimize_HideBacklash:      &Param{Value: "0"},
	Optimize_ReduceDirChanges:  &Param{Value: "0"},
	Optimize_ChooseCorners:     &Param{Value: "0"},
	Optimize_AllowReverse:      &Param{Value: "1"},
	Optimize_RemoveOverlaps:    &Param{Value: "0"},
	Optimize_OptimalEntryPoint: &Param{Value: "1"},
	Optimize_OverlapDist:       &Param{Value: "0.025"},
}
