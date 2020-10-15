package executor

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/kubernetes/cmd/kube-apiserver/app"
	cmapp "k8s.io/kubernetes/cmd/kube-controller-manager/app"
	proxy "k8s.io/kubernetes/cmd/kube-proxy/app"
	sapp "k8s.io/kubernetes/cmd/kube-scheduler/app"
	kubelet "k8s.io/kubernetes/cmd/kubelet/app"
)

func init() {
	executor = Embedded{}
}

type Embedded struct{}

func (Embedded) APIServer(ctx context.Context, etcdReady <-chan struct{}, args []string) (authenticator.Request, http.Handler, error) {
	<-etcdReady
	command := app.NewAPIServerCommand(ctx.Done())
	command.SetArgs(args)

	go func() {
		logrus.Fatalf("apiserver exited: %v", command.Execute())
	}()
	logrus.Info("<-app.StartupConfig")
	startupConfig := <-app.StartupConfig
	logrus.Info("<-app.StartupConfig done")
	return startupConfig.Authenticator, startupConfig.Handler, nil
}

func (Embedded) Scheduler(apiReady <-chan struct{}, args []string) error {
	command := sapp.NewSchedulerCommand()
	command.SetArgs(args)
	go func() {
		<-apiReady
		logrus.Fatalf("scheduler exited: %v", command.Execute())
	}()
	return nil
}

func (Embedded) ControllerManager(apiReady <-chan struct{}, args []string) error {
	command := cmapp.NewControllerManagerCommand()
	command.SetArgs(args)

	go func() {
		<-apiReady
		logrus.Fatalf("controller-manager exited: %v", command.Execute())
	}()

	return nil
}

func (Embedded) Kubelet(args []string) error {
	command := kubelet.NewKubeletCommand(context.Background())
	command.SetArgs(args)

	go func() {
		logrus.Fatalf("kubelet exited: %v", command.Execute())
	}()

	return nil
}

func (Embedded) KubeProxy(args []string) error {
	command := proxy.NewProxyCommand()
	command.SetArgs(args)

	go func() {
		logrus.Fatalf("kube-proxy exited: %v", command.Execute())
	}()

	return nil
}
