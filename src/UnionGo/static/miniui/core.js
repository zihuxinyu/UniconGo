/**
 * Created by weibaohui on 14-5-24.
 */


function GetParams(url, c) {
    if (!url) url = location.href;
    if (!c) c = "?";
    url = url.split(c)[1];
    var params = {};
    if (url) {
        var us = url.split("&");
        for (var i = 0, l = us.length; i < l; i++) {
            var ps = us[i].split("=");
            params[ps[0]] = decodeURIComponent(ps[1]);
        }
    }
    return params;
}


function onSkinChange(skin) {
    //mini.Cookie.set('miniuiSkin', skin);
    mini.Cookie.set('miniuiSkin', skin, 100);//100天过期的话，可以保持皮肤切换
    window.location.reload()
}
function AddCSSLink(id, url, doc) {
    doc = doc || document;
    var link = doc.createElement("link");
    link.id = id;
    link.setAttribute("rel", "stylesheet");
    link.setAttribute("type", "text/css");
    link.setAttribute("href", url);

    var heads = doc.getElementsByTagName("head");
    if (heads.length)
        heads[0].appendChild(link);
    else
        doc.documentElement.appendChild(link);
}
function alert(str) {
    mini.showTips({
        content: "<b>" + str + "</b> ",
        state: "warning",
        x: "center",
        y: "top",
        timeout: 1000
    });
}

/****************renderers*******************/
function onDateRenderer(e) {
    var date = mini.parseDate(e.value)
    var ss = mini.formatDate(date, "yyyy-MM-dd HH:mm:ss")
    return  ss;

}
/****************renderers*******************/

/***********grid op start***********/

function saveGrid(gridid, posturl) {
    var grid = mini.get(gridid);
    var data = grid.getChanges(null, true);
    //时间字段拿出来单独转换为beego需要的格式

    for (i = 0; i < data.length; i++) {
        if (data[i].Msgexpdate!=undefined) {
            data[i].Msgexpdate = mini.formatDate(data[i].Msgexpdate, "yyyy-MM-ddTHH:mm:ss+08:00");
        }
    }
    var json = mini.encode(data);
    alert(json);
    grid.loading("保存中，请稍后......");
    $.ajax({
        url: posturl,
        data: { data: json },
        type: "post",
        success: function (text) {
            grid.reload();
        },
        error: function (jqXHR, textStatus, errorThrown) {
            alert(jqXHR.responseText);
        }
    });
}
/***********grid op end***********/
