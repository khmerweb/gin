var episode = 0

const genJson = () => {
    const type = $('select[name="type"').val()
    const id = $('input[name="videoid"').val()
    const ending = $('select[name="status"').val()
            
    var entries = { type, id, ending }
        
    var success = false
    
    for(let v in entries){
        if(entries[v] === ''){
            alert('You need to fill the required field '+v)
            success = false
            break
        }else{
            success = true
        }
    }

    if(success){
        let json = $('input[name="videos"]').val()
        
        if(json === ''){
            json = JSON.stringify([entries])
            $('input[name="videos"]').val(json)
        }else{
            json = JSON.parse(json)
            json.push(entries)
            json = JSON.stringify(json)
            $('input[name="videos"').val(json)
        }

        let html = ``
        for(let v in entries){
            html += `<input value="${entries[v]}" />`
        }

        html += `<button title="Delete" onClick="deleteRow(event)" class="episode">${++episode}</button>`

        if($('.viddata .caption').html() === ''){
            const caption = `<div><b>ប្រភេទ​</b><b>អត្តសញ្ញាណ​</b><b>ចប់ឬ​នៅ?</b><b>ភាគ/លុប</b></div>`
            $('.viddata .caption').append(caption)
        }

        $('.viddata .part').prepend(`<div>${html}</div>`)
    }
}

function deleteRow(e) {
    e.target.parentElement.remove()
    
    let index = parseInt(e.target.innerHTML)
    index = index - 1
    let json = $('input[name="videos"]').val()
    json = JSON.parse(json)
    json.splice(index, 1);
    episode = json.length
    if(json.length === 0){
        json = ''
        $('.viddata .caption div').remove()
    }else{
        json = JSON.stringify(json)
    }
    $('input[name="videos"').val(json)
    counter = episode
    for(let v=0; v<episode; v++){
        $('.episode').eq(v).html(counter--)
    }
}

function submitForm(){
    let json = $('input[name="videos"]').val()
    if(json.length){
        let videos = JSON.parse(json)
        let newVideos = []
        let part = {}
        let key = {0:'type', 1:'id', 2:'status'}
        
        for(let v=0; v<videos.length; v++){
            for(let j=0; j<3; j++){
                part[key[j]] = $(`.viddata .part div:eq(${v}) input:eq(${j})`).val()
            }

            newVideos.push({...part})
        }
        
        let newJson = JSON.stringify(newVideos)
        $('input[name="videos"]').val(newJson)
    }
}