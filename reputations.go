package stackongo

import (
	"os"
	"strings"
	"fmt"
)

func (session Session) getReputations(path string, params map[string]string) (output *Reputations, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(Reputations))
	output = parsed_response.(*Reputations)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(output.Error_name + ": " + output.Error_message)
	}

	return
}

// ReputationChangesForUsers returns a subset of the reputation changes for users with given ids. 
func (session Session) ReputationChangesForUsers(ids []int, params map[string]string) (output *Reputations, error os.Error) {
	string_ids := []string{}
	for _, v := range ids {
		string_ids = append(string_ids, fmt.Sprintf("%v", v))
	}
	request_path := strings.Join([]string{"users", strings.Join(string_ids, ";"), "reputation"}, "/")
	return session.getReputations(request_path, params)
}
