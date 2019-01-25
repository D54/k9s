// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/k8sland/k9s/resource (interfaces: MetricsIfc)

package resource_test

import (
	k8s "github.com/k8sland/k9s/resource/k8s"
	pegomock "github.com/petergtz/pegomock"
	v1 "k8s.io/api/core/v1"
	"reflect"
	"time"
)

type MockMetricsIfc struct {
	fail func(message string, callerSkip ...int)
}

func NewMockMetricsIfc() *MockMetricsIfc {
	return &MockMetricsIfc{fail: pegomock.GlobalFailHandler}
}

func (mock *MockMetricsIfc) NodeMetrics() (k8s.Metric, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsIfc().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("NodeMetrics", params, []reflect.Type{reflect.TypeOf((*k8s.Metric)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 k8s.Metric
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(k8s.Metric)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockMetricsIfc) PerNodeMetrics(_param0 []v1.Node) (map[string]k8s.Metric, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsIfc().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PerNodeMetrics", params, []reflect.Type{reflect.TypeOf((*map[string]k8s.Metric)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 map[string]k8s.Metric
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(map[string]k8s.Metric)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockMetricsIfc) PodMetrics() (map[string]k8s.Metric, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsIfc().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PodMetrics", params, []reflect.Type{reflect.TypeOf((*map[string]k8s.Metric)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 map[string]k8s.Metric
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(map[string]k8s.Metric)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockMetricsIfc) VerifyWasCalledOnce() *VerifierMetricsIfc {
	return &VerifierMetricsIfc{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockMetricsIfc) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMetricsIfc {
	return &VerifierMetricsIfc{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockMetricsIfc) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMetricsIfc {
	return &VerifierMetricsIfc{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockMetricsIfc) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMetricsIfc {
	return &VerifierMetricsIfc{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMetricsIfc struct {
	mock                   *MockMetricsIfc
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMetricsIfc) NodeMetrics() *MetricsIfc_NodeMetrics_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "NodeMetrics", params, verifier.timeout)
	return &MetricsIfc_NodeMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MetricsIfc_NodeMetrics_OngoingVerification struct {
	mock              *MockMetricsIfc
	methodInvocations []pegomock.MethodInvocation
}

func (c *MetricsIfc_NodeMetrics_OngoingVerification) GetCapturedArguments() {
}

func (c *MetricsIfc_NodeMetrics_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMetricsIfc) PerNodeMetrics(_param0 []v1.Node) *MetricsIfc_PerNodeMetrics_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PerNodeMetrics", params, verifier.timeout)
	return &MetricsIfc_PerNodeMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MetricsIfc_PerNodeMetrics_OngoingVerification struct {
	mock              *MockMetricsIfc
	methodInvocations []pegomock.MethodInvocation
}

func (c *MetricsIfc_PerNodeMetrics_OngoingVerification) GetCapturedArguments() []v1.Node {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MetricsIfc_PerNodeMetrics_OngoingVerification) GetAllCapturedArguments() (_param0 [][]v1.Node) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]v1.Node, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.([]v1.Node)
		}
	}
	return
}

func (verifier *VerifierMetricsIfc) PodMetrics() *MetricsIfc_PodMetrics_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PodMetrics", params, verifier.timeout)
	return &MetricsIfc_PodMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MetricsIfc_PodMetrics_OngoingVerification struct {
	mock              *MockMetricsIfc
	methodInvocations []pegomock.MethodInvocation
}

func (c *MetricsIfc_PodMetrics_OngoingVerification) GetCapturedArguments() {
}

func (c *MetricsIfc_PodMetrics_OngoingVerification) GetAllCapturedArguments() {
}
