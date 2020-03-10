package databox

import "testing"

func TestValidateDataBoxJobName(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{
			input:    "",
			expected: false,
		},
		{
			input:    "hello",
			expected: true,
		},
		{
			input:    "_hello",
			expected: false,
		},
		{
			input:    "hello-",
			expected: false,
		},
		{
			input:    "hello!",
			expected: false,
		},
		{
			input:    "malcolm-in-the-middle",
			expected: true,
		},
		{
			input:    "hello.",
			expected: false,
		},
		{
			input:    "qwertyuioplkjhgfdsazxcv",
			expected: true,
		},
		{
			input:    "qwertyuioplkjhgfdsazxcva",
			expected: true,
		},
		{
			input:    "qwertyuioplkjhgfdsazxcvgg",
			expected: false,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q..", v.input)

		_, errors := ValidateDataBoxJobName(v.input, "name")
		actual := len(errors) == 0
		if v.expected != actual {
			t.Fatalf("Expected %t but got %t", v.expected, actual)
		}
	}
}

func TestValidateDataBoxJobContactName(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{
			input:    "",
			expected: false,
		},
		{
			input:    "hello",
			expected: true,
		},
		{
			input:    "_hello",
			expected: true,
		},
		{
			input:    "hello-",
			expected: true,
		},
		{
			input:    "hello!",
			expected: true,
		},
		{
			input:    "malcolm-in-the-middle",
			expected: true,
		},
		{
			input:    "hello.",
			expected: true,
		},
		{
			input:    "ahellodfasdfsdafsdfsdfsdfasdfsdfb",
			expected: true,
		},
		{
			input:    "ahellodfasdfsdafsdfsdfsdfasdfsdfbc",
			expected: true,
		},
		{
			input:    "ahellodfasdfsdafsdfsdfsdfasdfsdfbss",
			expected: false,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q..", v.input)

		_, errors := ValidateDataBoxJobContactName(v.input, "name")
		actual := len(errors) == 0
		if v.expected != actual {
			t.Fatalf("Expected %t but got %t", v.expected, actual)
		}
	}
}

func TestValidateDataBoxJobPhoneNumber(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{
			input:    "",
			expected: false,
		},
		{
			input:    "+1123456789",
			expected: true,
		},
		{
			input:    "123456789",
			expected: false,
		},
		{
			input:    "hello123",
			expected: false,
		},
		{
			input:    "hello!",
			expected: false,
		},
		{
			input:    "malcolm-in-the-middle",
			expected: false,
		},
		{
			input:    "hello.",
			expected: false,
		},
		{
			input:    "+1",
			expected: false,
		},
		{
			input:    "+12",
			expected: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q..", v.input)

		_, errors := ValidateDataBoxJobPhoneNumber(v.input, "phone_number")
		actual := len(errors) == 0
		if v.expected != actual {
			t.Fatalf("Expected %t but got %t", v.expected, actual)
		}
	}
}

func TestValidateDataBoxJobEmail(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{
			input:    "",
			expected: false,
		},
		{
			input:    "hello",
			expected: false,
		},
		{
			input:    "hello@microsoft.com",
			expected: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q..", v.input)

		_, errors := ValidateDataBoxJobEmail(v.input, "email")
		actual := len(errors) == 0
		if v.expected != actual {
			t.Fatalf("Expected %t but got %t", v.expected, actual)
		}
	}
}

func TestValidateDataBoxJobPhoneExtension(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{
			input:    "",
			expected: true,
		},
		{
			input:    "hello",
			expected: false,
		},
		{
			input:    "123",
			expected: true,
		},
		{
			input:    "1234",
			expected: true,
		},
		{
			input:    "12345",
			expected: false,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q..", v.input)

		_, errors := ValidateDataBoxJobPhoneExtension(v.input, "phone_extension")
		actual := len(errors) == 0
		if v.expected != actual {
			t.Fatalf("Expected %t but got %t", v.expected, actual)
		}
	}
}

func TestValidateDataBoxJobStreetAddress(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{
			input:    "",
			expected: false,
		},
		{
			input:    "16 TOWNSEND ST",
			expected: true,
		},
		{
			input:    "qwertyuiopasdasgfdhdghjjkljklzxcxbc",
			expected: true,
		},
		{
			input:    "qwertyuiopasdasgfdhdghjjkljklzxcxbcz",
			expected: false,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q..", v.input)

		_, errors := ValidateDataBoxJobStreetAddress(v.input, "street_address")
		actual := len(errors) == 0
		if v.expected != actual {
			t.Fatalf("Expected %t but got %t", v.expected, actual)
		}
	}
}

func TestValidateDataBoxJobPostCode(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{
			input:    "",
			expected: false,
		},
		{
			input:    "94107",
			expected: true,
		},
		{
			input:    "941079410794107",
			expected: false,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q..", v.input)

		_, errors := ValidateDataBoxJobPostCode(v.input, "post_code")
		actual := len(errors) == 0
		if v.expected != actual {
			t.Fatalf("Expected %t but got %t", v.expected, actual)
		}
	}
}

func TestValidateDataBoxJobCity(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{
			input:    "",
			expected: false,
		},
		{
			input:    "hello",
			expected: true,
		},
		{
			input:    "hellohellohellohellohellohellohello",
			expected: false,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q..", v.input)

		_, errors := ValidateDataBoxJobCity(v.input, "city")
		actual := len(errors) == 0
		if v.expected != actual {
			t.Fatalf("Expected %t but got %t", v.expected, actual)
		}
	}
}

func TestValidateDataBoxJobCompanyName(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{
			input:    "",
			expected: false,
		},
		{
			input:    "hello",
			expected: true,
		},
		{
			input:    "_hello",
			expected: true,
		},
		{
			input:    "hello-",
			expected: true,
		},
		{
			input:    "hello!",
			expected: true,
		},
		{
			input:    "malcolm-in-the-middle",
			expected: true,
		},
		{
			input:    "hello.",
			expected: true,
		},
		{
			input:    "qwertyuiopasdasgfdhdghjjkljklzxcxb",
			expected: true,
		},
		{
			input:    "qwertyuiopasdasgfdhdghjjkljklzxcxbc",
			expected: true,
		},
		{
			input:    "qwertyuiopasdasgfdhdghjjkljklzxcxbca",
			expected: false,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q..", v.input)

		_, errors := ValidateDataBoxJobCompanyName(v.input, "company_name")
		actual := len(errors) == 0
		if v.expected != actual {
			t.Fatalf("Expected %t but got %t", v.expected, actual)
		}
	}
}

func TestValidateDataBoxJobDiskPassKey(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{
			input:    "",
			expected: false,
		},
		{
			input:    "hellohellohello!1",
			expected: true,
		},
		{
			input:    "123123123123123!1",
			expected: true,
		},
		{
			input:    "2@hellohellohello",
			expected: true,
		},
		{
			input:    "2@hellohe5$llohello3&",
			expected: true,
		},
		{
			input:    "hellohellohello2",
			expected: false,
		},
		{
			input:    "hellohellohello#",
			expected: false,
		},
		{
			input:    "#hellohellohello",
			expected: false,
		},
		{
			input:    "2hellohellohello",
			expected: false,
		},
		{
			input:    "malcolm-in-the-middle",
			expected: false,
		},
		{
			input:    "malcolm1in2the3middle",
			expected: false,
		},
		{
			input:    "hellohellohello2",
			expected: false,
		},
		{
			input:    "bsadfasdfsdf2)bsadfasdfsdf2)acc",
			expected: true,
		},
		{
			input:    "bsadfasdfsdf2)bsadfasdfsdf2)accs",
			expected: true,
		},
		{
			input:    "bsadfasdfsdf2)bsadfasdfsdf2)accss",
			expected: false,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q..", v.input)

		_, errors := ValidateDataBoxJobDiskPassKey(v.input, "databox_disk_passkey")
		actual := len(errors) == 0
		if v.expected != actual {
			t.Fatalf("Expected %t but got %t", v.expected, actual)
		}
	}
}
