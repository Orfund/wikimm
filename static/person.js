let socIcon = (info)=>{
    let iconSrc;
    switch(info.network){
        case "istina":
            iconSrc = "/data/istina.png";
            break;
        case "youtube":
            iconSrc = "/data/youtube.png";
            break;
    }
    return `<img src="${iconSrc}" onclick="window.location.href='${info.link}'"/>`
};

let person = (pinfo)=>{
    return `<div class="person">
        <div class="pheader">
            <img src="${pinfo.ava}"/>
            <div class="pinfo">
                <div class="pname">${pinfo.name}</div>
                <div class="pplace">${pinfo.place}</div>
                <div class="links">${pinfo.social.map(socIcon)}</div>
        </div>
    </div>`
};
