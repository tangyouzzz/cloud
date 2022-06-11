axios.get("http://116.205.228.136:8848/user/index",{headers:{'authorization':window.localStorage.getItem("token")}}).then(function(response){
    // alert(response['data']['data']['username'])
    if (response["data"]["code"]==2000){
                login.log = 0;
                login.isLogin = true;
                login.username = response['data']['data']['username']
                alert(login.username+"  "+"huanyinghuilai")
            }else{
                login.log=1
            }
    
})
axios.get("http://116.205.228.136:8848/file/list",{headers:{'authorization':window.localStorage.getItem("token")}}).then(function(response){
                // item = response['data']['file'][1]
                file.tableData = response['data']['file']
})


var login = new Vue({
    el: "#login",
    data: {
        isLogin: false,
        log: 0,
        email: "",
        password: "",
        username:"",
    },
    methods:{
        login1: function () {
            this.log = 1
        },
        login2:function () {
            axios.post("http://116.205.228.136:8848/user/auth",{Username:this.email,Password:this.password}).then(function(response){
            // window.localStorage.setItem('token',response["data"]["data"]["token"])
            // var z =  window.localStorage.getItem("token")
            // alert(response["data"]["code"])
            if (response["data"]["code"]==2000){
                login.log = 0;
                login.isLogin = true;
                login.username = login.email
                window.location.reload()                        
                window.localStorage.setItem('token',response["data"]["data"]["token"])
                // alert("asf")
            }else{
                // console.log(response)
                alert(response["data"]["msg"])
            }
    })
            
        },
        close: function () {
            this.log = 0;
            //清空input中的内容
            this.email = "";
            this.password = "";
        },
        logout: function () {
            this.isLogin = false;
            this.email = "";
            this.password = "";
            window.localStorage.setItem('token',"")
            window.location.reload()
                                    
        },
    }
})
// var app = new Vue({
// el: '#app',
// data: {
//      email: '',
//      password:""
// },
// methods:{
//     yyy:function(){
//         // alert(this.email)
//         axios.post("http://116.205.228.136:8848/auth",{Username:this.email,Password:this.password}).then(function(response){
//         window.localStorage.setItem('token',response["data"]["data"]["token"])
//         var z =  window.localStorage.getItem("token")
//         console.log(z)
//         alert(z)
//     })
//     },
// }
// })
var file = new Vue({
    el:'#file',
    data:{
        tableData:[]
    },
    methods:{
        list:function(){
            axios.get("http://116.205.228.136:8848/file/list",{headers:{'authorization':window.localStorage.getItem("token")}}).then(function(response){
                if (response["data"]["code"]==2003||response["data"]["code"]==2005){
                    login.log=1}else{
                        file.tableData = response['data']['file']
                        // console.log(response['data']['file'])
                        // alert(response['data']['msg'])
                    }
        })
        },
        upload:function () {
            let file = document.getElementsByName('file')[0].files[0];
            let formData = new FormData();
            formData.append("file", file, file.name);

            const config = {
                headers: { "Content-Type": "multipart/form-data;boundary=" + new Date().getTime(),'authorization':window.localStorage.getItem("token") }
                 };
            axios.post("http://116.205.228.136:8848/file/upload", formData, config).then(function(response){
                if (response["data"]["code"]==2003||response["data"]["code"]==2005){
                    login.log=1
                }else{
                    alert(response["data"]["msg"])
                }
            })
            },

        download:function (row) {
            let fname =row["name"]
            axios.get("http://116.205.228.136:8848/file/download",{headers:{'authorization':window.localStorage.getItem("token")},params:{name:fname},'responseType':"blob"}).then(res=>{
                    let url = window.URL.createObjectURL(res.data)
                    let link =document.createElement("a")
                    link.style.display="none"
                    link.href=url
                    link.setAttribute("download",fname)
                    document.body.appendChild(link)
                    link.click()
                    document.body.removeChild(link)
                    window.URL.revokeObjectURL(link.href)
                    document.body.removeChild(link)

            }
                )
            },
        filedelete:function(row){
            let fname =row["name"]
            axios.get("http://116.205.228.136:8848/file/delete",{headers:{'authorization':window.localStorage.getItem("token")},params:{name:fname}}).then(function(response){
            if (response["data"]["code"]==2000){
                alert(response["data"]["msg"])
                // window.location.reload()
            }else{
                alert(response["data"]["msg"])
            }
            })
        }
    },
})  
