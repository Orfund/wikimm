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

let PersonName = (pinfo)=>(
  <div className="pinfo">
      <div>{pinfo.name}</div>
      <div>{pinfo.surname}</div>

  </div>
);

let Socials = (pinfo)=>(
    <div className="socwrapper">
        {console.log(pinfo.social)}
    </div>
);

let Person = (pinfo)=> (
    <div className="person">
        <div className="pheader">
            <img src={"data/"+pinfo.ava} alt=""/>
            <PersonName {...pinfo}/>
            <Socials {...pinfo}/>
        </div>
        <div>
            {pinfo.text}
        </div>
    </div>
    );
