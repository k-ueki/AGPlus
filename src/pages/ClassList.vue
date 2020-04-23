<template>
    <div class="hello">
<!--        <SideBar></SideBar>-->
        <div>
            <h2>授業一覧</h2>
        </div>
        <div class="card-board">
            <div class="card">
                <a>HogeHoge</a><br/>
                <a>教員名：hoge太郎</a>
            </div><br/>
            <div class="card">
                <a>piyopiyo</a><br/>
                <a>教員名：piyo太郎</a>
            </div><br/>
            <div class="card">
                <a>vowvow</a><br/>
                <a>教員名：vow太郎</a>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios'
    // import SideBar from './SideBar'

    const host = "localhost:7777"
    const getClassURL = host + "/classes"

    export default {
        name: 'ClassList',
        data(){
            return{
                classes:[],
            }
        },
        components:{
         // SideBar
        },
        props: {
            msg: String
        },
        created() {
            axios.get(getClassURL)
                .then(res=>{
                    this.classes = res
                })
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    h3 {
        margin: 40px 0 0;
    }
    ul {
        list-style-type: none;
        padding: 0;
    }
    li {
        display: inline-block;
        margin: 0 10px;
    }
    a {
        color: #42b983;
    }
    .card-board{

    }
    .card{
        width:10%;
        height:15%;
        background-color:black;
    }
</style>
