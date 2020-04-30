<template>
    <div class="hello">
<!--        <SideBar></SideBar>-->
        <div>
            <h2>授業一覧</h2>
        </div>
        <div class="card-board">
            <div v-for="(v,key) in classes" :key="key">
                <div class="card">
                    <a>{{v.Name}}</a><br/>
                    <a>教員名：{{v.Teacher}}</a><br/>
                </div><br/>
            </div>
        </div>
    </div>
</template>

<script>
    // import SideBar from './SideBar'
    import axios from 'axios'

    var baseURL = "http://localhost:8888"

    export default {
        name: 'ClassList',
        components:{
         // SideBar
        },
        props: {
            msg: String
        },
        data(){
            return{
                classes:[],
            }
        },
        created() {
            axios.get(baseURL + "/classes/")
                .then(response=>{
                    this.classes = response.data
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
