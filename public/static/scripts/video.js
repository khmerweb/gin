var episode = 0

function getCategory(){
    let categories = $('input[name="categories"').val()
    let category = $('select[name="category"').val()
    if(categories === ''){
        categories += category
    }else{
        categories += (`, ${category}`)
    }

    $('input[name="categories"').val(categories)
}

const genJson = () => {
    let videosStr = $('input[name="videos"]').val()
    if(videosStr.length){
        episode = (JSON.parse(videosStr)).length
    }
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

function loadVideos(rawVideos){
    if(rawVideos.length){
        const videos = JSON.parse(rawVideos)
        episode = videos.length
        for(let v=0; v<videos.length; v++){
            let html = ``
            for(let key in videos[v]){
                html += `<input value="${videos[v][key]}" />`
            }
            html += `<button title="Delete" onClick="deleteRow(event)" class="episode">${episode--}</button>`
            if($('.viddata .caption').html() === ''){
                const caption = `<div><b>ប្រភេទ​</b><b>អត្តសញ្ញាណ​</b><b>ចប់ឬ​នៅ?</b><b>ភាគ/លុប</b></div>`
                $('.viddata .caption').append(caption)
            }
            $('.viddata .part').append(`<div>${html}</div>`)
        }
    }
}

function addQueryString(){
    $('.editLink').on('click', function(event) {
        event.preventDefault();
        let currentHref = $(this).attr('href');
        let selectedPage = $('.pagination select').val();
        let newQueryString = `p=${selectedPage}`
        let newHref = currentHref + "?" + newQueryString;
        window.location.href = newHref
    })
}

$(document).ready(function(){
    const rawVideos = $('input[name="videos"]').val()
    loadVideos(rawVideos)
    addQueryString()
});



async function paginate(type, e){
	const response = await fetch(`/admin/${type}/paginate/${e.target.value}`)
	let Items = await response.json()
    $('.Items .items').empty()
    for(let item of Items.items){
        let html = `
        <div class="item">
            <a class="thumb" href="/${type}/${item.id}">
                <img src="${item.thumb}" />
                ${item.videos !== "" ? `<img class="play" src="/static/images/play.png" alt='' />` : ''}
            </a>
            <div class="title">
                <a href="/${type}"/${item.id}>${item.title}</a>
                <div>${new Date(item.date).toLocaleDateString('it-IT')}</div>
            </div>
            <div class="edit">
				<a href="/admin/${type}/delete/${item.id}" onclick="return confirm('Are you sure you want to delete this item?');" >
					<img src="/static/images/delete.png" alt=''/>
				</a>
                <a class="editLink" style="padding-right:5px;" href="/admin/${type}/edit/${item.id}">
					<img src="/static/images/edit.png" alt='' />
				</a>
            </div> 
        </div>
        `
        $('.Items .items').append(html)
        addQueryString()
    }
}