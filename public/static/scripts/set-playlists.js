
    
    let posts = playlists[0]
    let jq = $

    const dark = 'brightness(20%)'
    const normal = 'brightness(100%)'
    const laodingVideo = 'NcQQVbioeZk'
    
    let category = 'news'
    let pageAmount = Math.ceil(countPlaylists[0]/frontend)

    function parseVideos(posts){
        let videos = []
        let thumbs = []
        for(let post of posts){
            videos.push(JSON.parse(post.videos))
            thumbs.push(post.thumb)
        }
        videos.thumbs = thumbs
        return videos
    }

    let latestNews = parseVideos(playlists[0])
    latestNews.category = "news"
    let latestMovies = parseVideos(playlists[1])
    latestMovies.category = "movie"
    let latestTravel = parseVideos(playlists[2])
    latestTravel.category = "travel"
    let latestSimulation = parseVideos(playlists[3])
    latestSimulation.category = "simulation"
    let latestSport = parseVideos(playlists[4])
    latestSport.category = "sport"
    let latestDocumentary = parseVideos(playlists[5])
    latestDocumentary.category = "documentary"
    let latestFood = parseVideos(playlists[6])
    latestFood.category = "food"
    let latestMusic = parseVideos(playlists[7])
    latestMusic.category = "music"
    let latestGame = parseVideos(playlists[8])
    latestGame.category = "game"
    
    let rawPlaylist = {
        news: playlists[0],
        movie: playlists[1],
        travel: playlists[2],
        simulation: playlists[3],
        sport: playlists[4],
        documentary: playlists[5],
        food: playlists[6],
        music: playlists[7],
        game: playlists[8],
    }

    let playlistThumbs = {
        movie: rawPlaylist['movie'][0].thumb,
        travel: rawPlaylist['travel'][0].thumb,
        simulation: rawPlaylist['simulation'][0].thumb,
        sport: rawPlaylist['sport'][0].thumb,
        documentary: rawPlaylist['documentary'][0].thumb,
        food: rawPlaylist['food'][0].thumb,
        music: rawPlaylist['music'][0].thumb,
        game: rawPlaylist['game'][0].thumb,
    }

    let CountPlaylists = {
        news: countPlaylists[0],
        movie: countPlaylists[1],
        travel: countPlaylists[2],
        simulation: countPlaylists[3],
        sport: countPlaylists[4],
        documentary: countPlaylists[5],
        food: countPlaylists[6],
        music: countPlaylists[7],
        game: countPlaylists[8],
    }

    let videoPlaylists = {
        news: latestNews,
        movie: latestMovies,
        travel: latestTravel,
        simulation: latestSimulation,
        sport: latestSport,
        documentary: latestDocumentary,
        food: latestFood,
        music: latestMusic,
        game: latestGame,
    }

    $(document).ready(function() {
        $(`.random-video button #movie`).attr('src', playlistThumbs['movie'])
        $(`.random-video button #travel`).attr('src', playlistThumbs['travel'])
        $(`.random-video button #simulation`).attr('src', playlistThumbs['simulation'])
        $(`.random-video button #sport`).attr('src', playlistThumbs['sport'])
        $(`.random-video button #documentary`).attr('src', playlistThumbs['documentary'])
        $(`.random-video button #food`).attr('src', playlistThumbs['food'])
        $(`.random-video button #music`).attr('src', playlistThumbs['music'])
        $(`.random-video button #game`).attr('src', playlistThumbs['game'])

        $(`.random-video button:nth-child(1) p`).html(countPlaylists[1] + " ភាពយន្ត")
        $(`.random-video button:nth-child(2) p`).html(countPlaylists[2] + " ដើរលេង")
        $(`.random-video button:nth-child(3) p`).html(countPlaylists[3] + " ពិភព​និម្មិត")
        $(`.random-video button:nth-child(4) p`).html(countPlaylists[4] + " កីឡា")
        $(`.random-video button:nth-child(5) p`).html(countPlaylists[5] + " ឯកសារ")
        $(`.random-video button:nth-child(6) p`).html(countPlaylists[6] + " មុខ​ម្ហូប")
        $(`.random-video button:nth-child(7) p`).html(countPlaylists[7] + " របាំ​តន្ត្រី")
        $(`.random-video button:nth-child(8) p`).html(countPlaylists[8] + " ល្បែង​កំសាន្ត")
    });
    
    async function getRandomPlaylist(category, thumbs){
		const response = await fetch(`/api/playlist/${category}`, {
			method: 'POST',
			body: JSON.stringify({"thumbs": thumbs }),
			headers: {
				'Content-Type': 'application/json'
			}
		})
		const newPlaylist_ = await response.json()
        posts = newPlaylist_.playlist
        rawPlaylist[category] = newPlaylist_.playlist
        playlistThumbs[category] = newPlaylist_.playlist[0].thumb
        $(`.random-video button #${category}`).attr('src', newPlaylist_.playlist[0].thumb)
        let newPlaylist = parseVideos(newPlaylist_.playlist)
        newPlaylist.category = category
        videoPlaylists[category] = newPlaylist
        displayPosts(newPlaylist_.playlist, pageAmount)
        return newPlaylist
	}

    async function newPlaylist(){
        player.unMute()
        player.loadVideoById(laodingVideo)
        if(player.playlist.category !== 'news'){
            player.playlist = await getRandomPlaylist(player.playlist.category, player.playlist.thumbs) 
        }
        jq(`.Home .container .wrapper:nth-child(${player.part+1}) img`).css({'filter':normal})
        jq(`.Home .container .wrapper:nth-child(${player.part+1}) p`).css({'display':'none'})
        player.part = 0
        if(player.playlist[player.part][0].type === "YouTubePlaylist"){
            player.loadVideoById(initialVideoId)
            player.loadPlaylist({list:player.playlist[player.part][0].id,listType:'playlist',index:0})
        }else{
            player.index = 0
            if(!(player.playlist[player.part].reversal)){
                player.playlist[player.part].reverse()
                player.playlist[player.part].reversal = true
            }
            player.loadVideoById(player.playlist[player.part][0].id)
        }
        
        jq(`.Home .container .wrapper:nth-child(${player.part+1}) img`).css({'filter':dark})
        jq(`.Home .container .wrapper:nth-child(${player.part+1}) p`).css({'display':'block'})
    }

    function loadVideo(playlist){
        if(playlist[0][0].type === "YouTubePlaylist"){
            player.loadPlaylist({list:playlist[0][0].id,listType:'playlist',index:0})
        }else{
            playlist[0].reverse()
            playlist[0].reversal = true
            player.loadVideoById(playlist[0][0].id)
        }
        jq('.Home .container .wrapper:nth-child(1) img').css({'filter':dark})
        jq('.Home .container .wrapper:nth-child(1) p').css({'display':'block'})
    }

    function onPlayerReady(event) {
        player.part = 0
        player.index = 0
        player.thumb = 1
        player.label = 'ព័ត៌មាន'
        player.playlist = latestNews 
        displayPosts(rawPlaylist.news, pageAmount)
        loadVideo(latestNews)
    }

    function changeCategory(playlist, label, obj=0, thumb=0, part=0) {
        if(obj){posts = obj}
        if(label){player.label = label}
        if(playlist){player.playlist = playlist}
        
        category = player.playlist.category
        pageAmount = Math.ceil(CountPlaylists[player.playlist.category]/frontend)
        displayPosts(posts, pageAmount)

        if((player.playlist.category === 'home')||(player.playlist.category === 'news')){
            jq(`.random-video button:nth-child(${player.thumb}) img`).css({'filter':normal})
            jq(`.random-video button:nth-child(${player.thumb}) .playing`).css({'display':'none'})
        }
        if(thumb){
            jq(`.random-video button:nth-child(${player.thumb}) img`).css({'filter':normal})
            jq(`.random-video button:nth-child(${player.thumb}) .playing`).css({'display':'none'})
            player.thumb = thumb
            jq(`.random-video button:nth-child(${player.thumb}) img`).css({'filter':dark})
            jq(`.random-video button:nth-child(${player.thumb}) .playing`).css({'display':'block'})
        }
        jq(`.Home .container .wrapper:nth-child(${player.part+1}) img`).css({'filter':normal})
        jq(`.Home .container .wrapper:nth-child(${player.part+1}) p`).css({'display':'none'})
        player.part = part
        player.unMute()
        if(player.playlist[player.part][0].type === "YouTubePlaylist"){
            player.loadVideoById(initialVideoId)
            player.loadPlaylist({list:player.playlist[player.part][0].id,listType:'playlist',index:0})
            jq('.latest-video').html(player.label)
        }else{
            if(!(player.playlist[player.part].reversal)){
                player.playlist[player.part].reverse()
                player.playlist[player.part].reversal = true
            }
            player.loadVideoById(player.playlist[player.part][0].id)
            jq('.latest-video').html(player.label)
            
        }
        jq(`.Home .container .wrapper:nth-child(${player.part+1}) img`).css({'filter':dark})
        jq(`.Home .container .wrapper:nth-child(${player.part+1}) p`).css({'display':'block'})
    }

    function onPlayerError(event){
        if(player.index + 1 < player.playlist[player.part].length){
            player.index += 1
            player.loadVideoById(player.playlist[player.part][player.index].id)
        }else{
            jq(`.Home .container .wrapper:nth-child(${player.part+1}) img`).css({'filter':normal})
            jq(`.Home .container .wrapper:nth-child(${player.part+1}) p`).css({'display':'none'})
            player.part += 1
            if(player.part === player.playlist.length){
                player.part = 0
            }

            if(player.playlist[player.part][0].type === "YouTubePlaylist"){
                player.loadVideoById(initialVideoId)
                player.loadPlaylist({list:player.playlist[player.part][0].id,listType:'playlist',index:0})
            }else{
                player.index = 0
                if(!(player.playlist[player.part].reversal)){
                    player.playlist[player.part].reverse()
                    player.playlist[player.part].reversal = true
                }
                player.loadVideoById(player.playlist[player.part][0].id)
            }
            jq(`.Home .container .wrapper:nth-child(${player.part+1}) img`).css({'filter':dark})
            jq(`.Home .container .wrapper:nth-child(${player.part+1}) p`).css({'display':'block'})
        }
    }

    async function onPlayerStateChange(event) {   
        if(event.data === YT.PlayerState.ENDED){
            if(player.index + 1 < player.playlist[player.part].length){
                player.index += 1
                player.loadVideoById(player.playlist[player.part][player.index].id)
                
            }else{
                jq(`.Home .container .wrapper:nth-child(${player.part+1}) img`).css({'filter':normal})
                jq(`.Home .container .wrapper:nth-child(${player.part+1}) p`).css({'display':'none'})
                player.part += 1
                if(player.part === player.playlist.length){
                    player.loadVideoById(laodingVideo)
                    if(player.playlist.category !== 'news'){
                        player.playlist = await getRandomPlaylist(player.playlist.category, player.playlist.thumbs)
                    }
                    player.part = 0
                }

                if(player.playlist[player.part][0].type === "YouTubePlaylist"){
                    player.loadVideoById(initialVideoId)
                    player.loadPlaylist({list:player.playlist[player.part][0].id,listType:'playlist',index:0})
                }else{
                    player.index = 0
                    if(!(player.playlist[player.part].reversal)){
                        player.playlist[player.part].reverse()
                        player.playlist[player.part].reversal = true
                    }
                    player.loadVideoById(player.playlist[player.part][0].id)
                }
                jq(`.Home .container .wrapper:nth-child(${player.part+1}) img`).css({'filter':dark})
                jq(`.Home .container .wrapper:nth-child(${player.part+1}) p`).css({'display':'block'})
            }
        }
    }

    function nextPrevious(move){
        player.unMute()
        if(move === "next"){
            if(player.index + 1 < player.playlist[player.part].length){
                player.index += 1
                player.loadVideoById(player.playlist[player.part][player.index].id)
            }else{
                jq(`.Home .container .wrapper:nth-child(${player.part+1}) img`).css({'filter':normal})
                jq(`.Home .container .wrapper:nth-child(${player.part+1}) p`).css({'display':'none'})
                player.part += 1
                if(player.part === player.playlist.length){
                    player.part = 0
                }

                if(player.playlist[player.part][0].type === "YouTubePlaylist"){
                    player.loadVideoById(initialVideoId)
                    player.loadPlaylist({list:player.playlist[player.part][0].id,listType:'playlist',index:0})
                }else{
                    player.index = 0
                    if(!(player.playlist[player.part].reversal)){
                        player.playlist[player.part].reverse()
                        player.playlist[player.part].reversal = true
                    }
                    player.loadVideoById(player.playlist[player.part][0].id)
                }
                jq(`.Home .container .wrapper:nth-child(${player.part+1}) img`).css({'filter':dark})
                jq(`.Home .container .wrapper:nth-child(${player.part+1}) p`).css({'display':'block'})
            }
        }else if(move === "previous"){
            if(player.index > 0){
                player.index -= 1
                player.loadVideoById(player.playlist[player.part][player.index].id)
            }else{
                jq(`.Home .container .wrapper:nth-child(${player.part+1}) img`).css({'filter':normal})
                jq(`.Home .container .wrapper:nth-child(${player.part+1}) p`).css({'display':'none'})
                player.part -= 1
                if(player.part < 0){
                    player.part = player.playlist.length - 1
                }

                if(player.playlist[player.part][0].type === "YouTubePlaylist"){
                    player.loadVideoById(initialVideoId)
                    player.loadPlaylist({list:player.playlist[player.part][0].id,listType:'playlist',index:0})
                }else{
                    player.index = 0
                    if(!(player.playlist[player.part].reversal)){
                        player.playlist[player.part].reverse()
                        player.playlist[player.part].reversal = true
                    }
                    player.loadVideoById(player.playlist[player.part][0].id)
                }
                jq(`.Home .container .wrapper:nth-child(${player.part+1}) img`).css({'filter':dark})
                jq(`.Home .container .wrapper:nth-child(${player.part+1}) p`).css({'display':'block'})
            }
        }
    }

    const ytPlayerId = 'youtube-player'
    let initialVideoId = 'cdwal5Kw3Fc'
    let player

    
    function load() {
        player = new YT.Player(ytPlayerId, {
            height: '390',
            width: '640',
            videoId: initialVideoId,
            playerVars: {
                'playsinline': 1,
                "enablejsapi": 1,
                "mute": 1,
                "autoplay": 1,
                "rel": 0,
            },
            events: {
                'onReady': onPlayerReady,
                'onStateChange': onPlayerStateChange,
                'onError': onPlayerError
            }
        })
    }
        
var tag = document.createElement('script');
tag.src = "https://www.youtube.com/iframe_api";
var firstScriptTag = document.getElementsByTagName('script')[0];
firstScriptTag.parentNode.insertBefore(tag, firstScriptTag);

function onYouTubeIframeAPIReady() {
    load()
} 

function displayPosts(posts, pageAmount){
    let html = `<div class="container">`
    for(const index in posts){
        html += `<div class="wrapper">`
        html += `<a onclick="changeCategory(false, false, false, false, ${index})">`
        html += `<img src="${posts[index].thumb}" alt=''/>`
        if(posts[index].videos.length){
            html +=`<img class="play-icon" src="/static/images/play.png" alt=''/>`
        }
        html += `<p>កំពុង​លេង...</p>`
        html += `</a>`
        html += `<div class="date">${(new Date(posts[index].date)).toLocaleDateString('it-IT')}</div>`
        html += `<a class="title" onclick="changeCategory(false, false, false, false, ${index})">`
        html += `<div >${posts[index].title}</div>`
        html += `</a>`
        html += `</div>`
    }
    html += `</div>`
    html += displayHomePagination(pageAmount)
    $(".Home").empty()
    $(".Home").html(html)
}

function paginate(event){
    document.location = "/" + category + "/" + event.target.value
}

function displayHomePagination(pageAmount){
    let footer = `<div class="pagination">`
    footer += `<span>​​​​​​​​​​​​​​​​​​​​​ទំព័រ </span> `
    footer += `<select onchange="paginate(event)">`
	for(const page of [...Array(pageAmount).keys()]){
		footer += `<option>${page+1}</option>`
	}
	footer += `</select>` 
    footer += ` <span>នៃ ${pageAmount}</span>`   
    footer += `</div> `
    return footer
}