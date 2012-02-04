package stackongo

import (
	"os"
	"strings"
)

func (session Session) getTagWikis(path string, params map[string]string) (output *TagWikis, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(TagWikis))
	output = parsed_response.(*TagWikis)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(output.Error_name + ": " + output.Error_message)
	}

	return

}

// WikisForTags returns the wikis that go with the given set of tags 
func (session Session) WikisForTags(tags []string, params map[string]string) (output *TagWikis, error os.Error) {
	request_path := strings.Join([]string{"tags", strings.Join(tags, ";"), "wikis"}, "/")
	return session.getTagWikis(request_path, params)
}
