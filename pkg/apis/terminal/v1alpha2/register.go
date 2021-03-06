/*

 Copyright 2019 The KubeSphere Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.

*/
package v1alpha2

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubesphere.io/kubesphere/pkg/apiserver/runtime"
	"kubesphere.io/kubesphere/pkg/apiserver/terminal"
	"kubesphere.io/kubesphere/pkg/models"
)

const GroupName = "terminal.kubesphere.io"

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha2"}

var (
	WebServiceBuilder = runtime.NewContainerBuilder(addWebService)
	AddToContainer    = WebServiceBuilder.AddToContainer
)

func addWebService(c *restful.Container) error {

	webservice := runtime.NewWebService(GroupVersion)

	tags := []string{"Terminal"}

	webservice.Route(webservice.GET("/namespaces/{namespace}/pods/{pod}").
		To(terminal.CreateTerminalSession).
		Doc("create terminal session").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(models.PodInfo{}))

	path := runtime.ApiRootPath + "/" + GroupVersion.String() + "/sockjs"
	c.Handle(path+"/", terminal.NewTerminalHandler(path))

	c.Add(webservice)

	return nil
}
