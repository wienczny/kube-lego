// Automatically generated by MockGen. DO NOT EDIT!
// Source: ../kubelego_const/interfaces.go

package mocks

import (
	logrus "github.com/Sirupsen/logrus"
	gomock "github.com/golang/mock/gomock"
	kubelego "github.com/jetstack/kube-lego/pkg/kubelego_const"
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/client/unversioned"
)

// Mock of KubeLego interface
type MockKubeLego struct {
	ctrl     *gomock.Controller
	recorder *_MockKubeLegoRecorder
}

// Recorder for MockKubeLego (not exported)
type _MockKubeLegoRecorder struct {
	mock *MockKubeLego
}

func NewMockKubeLego(ctrl *gomock.Controller) *MockKubeLego {
	mock := &MockKubeLego{ctrl: ctrl}
	mock.recorder = &_MockKubeLegoRecorder{mock}
	return mock
}

func (_m *MockKubeLego) EXPECT() *_MockKubeLegoRecorder {
	return _m.recorder
}

func (_m *MockKubeLego) KubeClient() *unversioned.Client {
	ret := _m.ctrl.Call(_m, "KubeClient")
	ret0, _ := ret[0].(*unversioned.Client)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) KubeClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KubeClient")
}

func (_m *MockKubeLego) Log() *logrus.Entry {
	ret := _m.ctrl.Call(_m, "Log")
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) Log() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Log")
}

func (_m *MockKubeLego) AcmeClient() kubelego.Acme {
	ret := _m.ctrl.Call(_m, "AcmeClient")
	ret0, _ := ret[0].(kubelego.Acme)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) AcmeClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AcmeClient")
}

func (_m *MockKubeLego) LegoHTTPPort() string {
	ret := _m.ctrl.Call(_m, "LegoHTTPPort")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoHTTPPort() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoHTTPPort")
}

func (_m *MockKubeLego) LegoEmail() string {
	ret := _m.ctrl.Call(_m, "LegoEmail")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoEmail() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoEmail")
}

func (_m *MockKubeLego) LegoURL() string {
	ret := _m.ctrl.Call(_m, "LegoURL")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) LegoURL() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LegoURL")
}

func (_m *MockKubeLego) Version() string {
	ret := _m.ctrl.Call(_m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) Version() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Version")
}

func (_m *MockKubeLego) AcmeUser() (map[string][]byte, error) {
	ret := _m.ctrl.Call(_m, "AcmeUser")
	ret0, _ := ret[0].(map[string][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKubeLegoRecorder) AcmeUser() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AcmeUser")
}

func (_m *MockKubeLego) SaveAcmeUser(_param0 map[string][]byte) error {
	ret := _m.ctrl.Call(_m, "SaveAcmeUser", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockKubeLegoRecorder) SaveAcmeUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveAcmeUser", arg0)
}

// Mock of Acme interface
type MockAcme struct {
	ctrl     *gomock.Controller
	recorder *_MockAcmeRecorder
}

// Recorder for MockAcme (not exported)
type _MockAcmeRecorder struct {
	mock *MockAcme
}

func NewMockAcme(ctrl *gomock.Controller) *MockAcme {
	mock := &MockAcme{ctrl: ctrl}
	mock.recorder = &_MockAcmeRecorder{mock}
	return mock
}

func (_m *MockAcme) EXPECT() *_MockAcmeRecorder {
	return _m.recorder
}

func (_m *MockAcme) ObtainCertificate(domains []string) (map[string][]byte, error) {
	ret := _m.ctrl.Call(_m, "ObtainCertificate", domains)
	ret0, _ := ret[0].(map[string][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAcmeRecorder) ObtainCertificate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ObtainCertificate", arg0)
}

// Mock of Tls interface
type MockTls struct {
	ctrl     *gomock.Controller
	recorder *_MockTlsRecorder
}

// Recorder for MockTls (not exported)
type _MockTlsRecorder struct {
	mock *MockTls
}

func NewMockTls(ctrl *gomock.Controller) *MockTls {
	mock := &MockTls{ctrl: ctrl}
	mock.recorder = &_MockTlsRecorder{mock}
	return mock
}

func (_m *MockTls) EXPECT() *_MockTlsRecorder {
	return _m.recorder
}

func (_m *MockTls) Hosts() []string {
	ret := _m.ctrl.Call(_m, "Hosts")
	ret0, _ := ret[0].([]string)
	return ret0
}

func (_mr *_MockTlsRecorder) Hosts() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Hosts")
}

func (_m *MockTls) SecretMetadata() *api.ObjectMeta {
	ret := _m.ctrl.Call(_m, "SecretMetadata")
	ret0, _ := ret[0].(*api.ObjectMeta)
	return ret0
}

func (_mr *_MockTlsRecorder) SecretMetadata() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SecretMetadata")
}

func (_m *MockTls) IngressMetadata() *api.ObjectMeta {
	ret := _m.ctrl.Call(_m, "IngressMetadata")
	ret0, _ := ret[0].(*api.ObjectMeta)
	return ret0
}

func (_mr *_MockTlsRecorder) IngressMetadata() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IngressMetadata")
}

func (_m *MockTls) Process() error {
	ret := _m.ctrl.Call(_m, "Process")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTlsRecorder) Process() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Process")
}

// Mock of Ingress interface
type MockIngress struct {
	ctrl     *gomock.Controller
	recorder *_MockIngressRecorder
}

// Recorder for MockIngress (not exported)
type _MockIngressRecorder struct {
	mock *MockIngress
}

func NewMockIngress(ctrl *gomock.Controller) *MockIngress {
	mock := &MockIngress{ctrl: ctrl}
	mock.recorder = &_MockIngressRecorder{mock}
	return mock
}

func (_m *MockIngress) EXPECT() *_MockIngressRecorder {
	return _m.recorder
}

func (_m *MockIngress) Tls() []kubelego.Tls {
	ret := _m.ctrl.Call(_m, "Tls")
	ret0, _ := ret[0].([]kubelego.Tls)
	return ret0
}

func (_mr *_MockIngressRecorder) Tls() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Tls")
}

func (_m *MockIngress) Ignore() bool {
	ret := _m.ctrl.Call(_m, "Ignore")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockIngressRecorder) Ignore() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ignore")
}
