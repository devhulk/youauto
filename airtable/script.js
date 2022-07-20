//output.markdown('# New Video Idea');

//let name = await input.textAsync('What is the name of the video?');
//let sponsor = await input.textAsync('What is the name of the video?');
//let date = 
//output.text(`Generating "${name}" assets in Notion and Dropbox.`);


// Fetch records from Notion.

let request_body = {
    autorename: false,
    force_async: false,
    paths: ["/YouTube/Airtable/a","/YouTube/Airtable/b","/YouTube/Airtable/graphics"]
}

let response = await fetch('https://api.dropboxapi.com/2/files/create_folder_batch', {
    method: 'POST', 
    body: JSON.stringify(request_body),
    headers: {
        "Authorization": "Bearer " + process.env.DROPBOX_TOKEN,
        "Content-Type": "application/json"
    }
});

console.log(await response.text());
