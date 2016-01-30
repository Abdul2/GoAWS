package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"net/http"
	"html/template"
//	"sort"
	"os"
)


//load templates first
var templ =	template.Must(template.ParseGlob("./templates/*"))

//construct to define fields to hold ec2 attributes
type Instance struct {

	InstanceId string
	InstanceType string
	VpcId string
	PrivateIpAddress string
	Tag string
	Architecture string
	Counter int
}


//slice to hold the result
type listofinstances []Instance


//implement the sort interface. this allow listofinstances to use sort
func (slice listofinstances) Len() int{

	return len(slice)
}

func (slice listofinstances) Less(i,j int) bool{

	return slice[i].PrivateIpAddress < slice[j].PrivateIpAddress;

}

func (slice listofinstances) Swap(i,j int){

	slice[i], slice[j] = slice[j], slice[j]
}


//diplay result
func awsinfohandler(w http.ResponseWriter, r *http.Request) {


	templ.ExecuteTemplate(w,"aws", getinstances())


}

func indexhandler(w http.ResponseWriter, r *http.Request){

	templ.ExecuteTemplate(w,"index",nil)
}

//interate over theh result and return a slice containingg the result
func getinstances() []Instance {


	var h Instance

	var listofcurrent listofinstances

	var tag string


	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")})

	resp, err := svc.DescribeInstances(nil)


	if err != nil {

		panic(err)

	}



	for idx, res := range resp.Reservations {

		fmt.Println("  > Number of Reservations: ", len(res.Instances))

		for _, inst := range resp.Reservations[idx].Instances {



			h.InstanceId = fmt.Sprintf("%s", *inst.InstanceId)
			h.InstanceType = fmt.Sprintf("%s", *inst.InstanceType)
			h.VpcId = fmt.Sprintf("%s", *inst.VpcId)
			h.PrivateIpAddress = fmt.Sprintf("%s", *inst.PrivateIpAddress)
			h.Architecture = fmt.Sprintf("%s", *inst.Architecture)
			h.Counter++


			for _, keys := range inst.Tags {
				if *keys.Key == "Name" {
					tag = *keys.Value
				}
			}

			h.Tag = tag

			listofcurrent = append(listofcurrent,h)
		}

	}


//	sort.Sort(listofcurrent)

	return listofcurrent
}


func GetPort() string {

	p := os.Getenv("PORT")

	if p != "" {

		return ":" + p
	}

	return ":8080"
}




func main() {


	http.HandleFunc("/instances", awsinfohandler)

	http.HandleFunc("/",indexhandler)

	http.ListenAndServe(GetPort(),nil)




}
