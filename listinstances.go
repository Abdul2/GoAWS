package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"net/http"
	"html/template"
	"sort"
)


//load templates first
var templ =	template.Must(template.ParseGlob("./templates/*"))

//construct to define fields to hold ec2 attributes
type HMPOinstance struct {

	InstanceId string
	InstanceType string
	VpcId string
	PrivateIpAddress string
	Tag string
	Counter int
}


//slice to hold the result
type listofHMPOinstances []HMPOinstance


//implement the sort interface. this allow lisofHMPOintsances to use sort
func (slice listofHMPOinstances) Len() int{

	return len(slice)
}

func (slice listofHMPOinstances) Less(i,j int) bool{

	return slice[i].PrivateIpAddress < slice[j].PrivateIpAddress;

}

func (slice listofHMPOinstances) Swap(i,j int){

	slice[i], slice[j] = slice[j], slice[j]
}


//diplay result
func awsinfohandler(w http.ResponseWriter, r *http.Request) {


	templ.ExecuteTemplate(w,"aws", HMPOInstances())


}


//interate over theh result and return a slice containingg the result
func HMPOInstances() []HMPOinstance{


	var h HMPOinstance

	var listofcurrent listofHMPOinstances

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


	sort.Sort(listofcurrent)

	return listofcurrent
}



func main() {


	http.HandleFunc("/", awsinfohandler)

	http.ListenAndServe(":8080",nil)




}
