// API Reference: https://www.wix.com/velo/reference/api-overview/introduction
// “Hello, World!” Example: https://learn-code.wix.com/en/article/1-hello-world
import {getJobs} from 'backend/Modules/Jobs/getJobs';

import wixUsers from 'wix-users';
import wixWindow from 'wix-window';
import {local} from 'wix-storage';

$w.onReady(function () {
	// admin user perms
	setTypeOfInput();
	wixUsers.currentUser.getRoles()
	.then((roles) => {
		const currentUser = wixUsers.currentUser;
		if(currentUser.role === "Admin"){
			$w('#button16').hide();
		} else {
			$w('#button16').show();
		}
	});

	// set jobsRepeater callback
	$w("#jobsRepeater").onItemReady( ($item, itemData, index) => {
    		$item("#jobTitle").text = itemData.title;
    		$item("#jobDescription").text = itemData.description;
			$item("#jobLocation").text = itemData.location;
			$item("#labels").text = itemData.labels.toString();
			$item("#jobId").text = itemData._id;
  	});

	refreshJobs()
});

$w("#jobsPagination").currentPage = 1;

export function refreshJobs() {
	// fetch jobs into repeater
	getJobs().then(jobInfo => {
    	$w("#jobsRepeater").data = []
		$w("#jobsRepeater").data = jobInfo;
	})
}

/**
*	Adds an event handler that runs when the element is clicked.
	[Read more](https://www.wix.com/corvid/reference/$w.ClickableMixin.html#onClick)
*	 @param {$w.MouseEvent} event
*/
export function deleteJob_click(event) {
	// This function was added from the Properties & Events panel. To learn more, visit http://wix.to/UcBnC-4
	// Add your code for this event here: 
	local.setItem("jobToDelete", $w("#jobId").text)
	wixWindow.openLightbox("DeleteJob").then(() => {
		refreshJobs();	
	});
}

/**
*	Adds an event handler that runs when the element is clicked.
	[Read more](https://www.wix.com/corvid/reference/$w.ClickableMixin.html#onClick)
*	 @param {$w.MouseEvent} event
*/
export function AddJobButton_click(event) {
	// This function was added from the Properties & Events panel. To learn more, visit http://wix.to/UcBnC-4
	// Add your code for this event here: 
	wixWindow.openLightbox("AddJob").then(() => {
		refreshJobs();
	});
}

/**
*	Adds an event handler that runs when the element is clicked.
	[Read more](https://www.wix.com/corvid/reference/$w.ClickableMixin.html#onClick)
*	 @param {$w.MouseEvent} event
*/
export function searchJobsButton_click(event) {
	// This function was added from the Properties & Events panel. To learn more, visit http://wix.to/UcBnC-4
	// Add your code for this event here: 
	getJobs($w("#searchJobInput").value).then(jobInfo => {
    	$w("#jobsRepeater").data = []
		$w("#jobsRepeater").data = jobInfo;
	})
}

function setTypeOfInput(){ // Insert in init of page
  $w("#firstNameInput").inputType = "text";
  $w("#lastNameInput").inputType = "text";
  $w("#emailInput").inputType = "email";
  $w("#phoneNumberInput").inputType = "tel";
  $w("#firstNameInput").required = true;
  $w("#lastNameInput").required = true;
  $w("#emailInput").required = true;
  $w("#phoneNumberInput").required = true;
  $w("#dateOfBirthPicker").required = true;
  $w("#jobDropDown").required = true;
  $w("#uploadPdfButton").required = true;
  $w("#uploadPdfButton").fileType = "Document";
}


export function postApplicant(firstName, lastName, email, dateOfBirth, phoneNumber, job, cv){
  const url = 'https://ezrecruit-backend-ryo2vcvbqq-uc.a.run.app/applicants';
  return fetch(url, {
    method: 'POST',
    headers: {
 'Content-Type': 'application/json'
        },
    body: JSON.stringify({
            firstName: firstName,
			lastName: lastName,
			email: email,
			phoneNumber: phoneNumber,
      dateOfBirth: dateOfBirth,
			job: job
            })    
    }).then( (httpResponse) => {
 if (httpResponse.ok) {
 return httpResponse.json();
        } else {
 return Promise.reject("Fetch did not succeed");
        }
    } );
}

/**
*	Adds an event handler that runs when the element is clicked.
	[Read more](https://www.wix.com/corvid/reference/$w.ClickableMixin.html#onClick)
*	 @param {$w.MouseEvent} event
*/
export function button1_click_1(event) {
	// This function was added from the Properties & Events panel. To learn more, visit http://wix.to/UcBnC-4
	// Add your code for this event here: 
  if(!$w("#firstNameInput").valid){
    console.log("first name not valid")
    return;
  }
  if(!$w("#lastNameInput").valid){
    console.log("last name not valid")
    return;
  }
  if(!$w("#emailInput").valid){
    console.log("email not valid")
    return;
  }
  if(!$w("#phoneNumberInput").valid){
    console.log("phone number not valid")
    return;
  }
  if(!$w("#dateOfBirthPicker").valid){
    console.log("date of birth not valid")
    return;
  }
  if(!$w("#jobDropDown").valid){
    console.log("job not valid")
    return;
  }
  // we are all good to go
  let firstName = $w("#firstNameInput").value;
  let lastName = $w("#lastNameInput").value;
  let email = $w("#emailInput").value;
  let phone = $w("#phoneNumberInput").value;
  let dateOfBirth = $w("#dateOfBirthPicker").value;
  let job = $w("#jobDropDown").value;
  let cv;
  $w("#uploadPdfButton").uploadFiles().then(uploadedFile => {
    console.log("Upload successful.");
    cv = uploadedFile;
  }).catch((uploadError) => {
    console.log("File upload error: " + uploadError.errorCode);
    console.log(uploadError.errorDescription);
    return;
    });
	console.log("and now for post applicant!!");
  postApplicant(firstName,lastName,email,dateOfBirth,phone,job,cv);
}