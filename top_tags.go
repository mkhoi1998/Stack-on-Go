package stackongo

import (
	"os"
	"strings"
	"fmt"
)

func (session Session) getTopTags(path string, params map[string]string) (output *TopTags, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(TopTags))
	output = parsed_response.(*TopTags)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(output.Error_name + ": " + output.Error_message)
	}

	return

}

// TopTagsByAnswerForUser returns a single user's top tags by answer score. 
func (session Session) TopTagsByAnswerForUser(id int, params map[string]string) (output *TopTags, error os.Error) {
	request_path := strings.Join([]string{"users", fmt.Sprintf("%v", id), "top-answer-tags"}, "/")
	return session.getTopTags(request_path, params)
}

// TopTagsByQuestionForUser returns a single user's top tags by question score. 
func (session Session) TopTagsByQuestionForUser(id int, params map[string]string) (output *TopTags, error os.Error) {
	request_path := strings.Join([]string{"users", fmt.Sprintf("%v", id), "top-question-tags"}, "/")
	return session.getTopTags(request_path, params)
}
