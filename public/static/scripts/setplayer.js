//public/scripts/setPlayer.js

function setScreen(video,id,click){
    if(video['vidType'] == 'OK'){
        var url = `//ok.ru/videoembed/${video['id']}`
    }else if(video['vidType'] == 'YouTube'){
        var url = `https://www.youtube.com/embed/${video['id']}`
    }else if(video['vidType'] == 'YouTubePlaylist'){
        var url = `https://www.youtube.com/embed/videoseries?list=${video['id']}`
    }else if(video['vidType'] === "Facebook"){
        var url = `https://www.facebook.com/watch/?v=${video['id']}`
    }else if(video['vidType'] === "GoogleDrive"){
        var url = `https://docs.google.com/file/d/${video['id']}/preview`
    }else if(video['vidType'] === "Vimeo"){
        var url = `https://player.vimeo.com/video/${video['id']}`
    }else if(video['vidType'] === "Dailymotion"){
        var url = `https://www.dailymotion.com/embed/video/${video['id']}`
    }

    if(video['vidType'] !== 'Facebook'){
        var iframe = `<div><iframe  src="${url}" frameborder="0" allowfullscreen></iframe></div>`
      }else{
        var iframe = `<div class="fb-video" data-href="${url}" data-width="auto" data-show-captions="true"></div>`
    }

    if(click){
        $('.Post .video .playlist #part'+clicked)
            .css({'background':'var(--playlist-part)', 'color': 'var(--playlist-color)'})
    }
    
    $('.Post .video .playlist #part'+id)
        .css({'background':'var(--playlist-clicked-bg)', 'color': 'var(--playlist-clicked)'})

    $('.screen').html(iframe) 

    clicked = id
}