// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/openshift/console-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Consoles returns a ConsoleInformer.
	Consoles() ConsoleInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Consoles returns a ConsoleInformer.
func (v *version) Consoles() ConsoleInformer {
	return &consoleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}