# go-reconcile

Super tiny go package which does reconcile planning: taking desired list and current observed list and returning the action to take to reach desired states (the items adding to or deleting from the current list). 

This can be used, for example, when you want to enforce only desired members to have some cloud IAM role. You can ask your cloud provider API about the current members which has the target roles for now. Then you can prepare the desired list of members to have the role and, by reconcile planning, you can get which member to add to or delete. Based on this, you can call cloud provider API and reach to the desired state.

I wrote this package to play with Go1.18 generics functinality and you should not use this in your code. I recommend you to write your own if you need similar functionality.

*NOTE*: To run this code, you need to use Go1.18 or later

Documentation is available on GoDoc: https://pkg.go.dev/github.com/tcnksm/go-reconcile