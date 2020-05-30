<template>
    <div class="wrapper" align="center">
        <h1 class="text-center" style="height:60px">Classes</h1>
        <v-divider></v-divider>
        {{ items }}
        <div v-for="(item, key, index) in items" :key="index">
            <template>
                  <v-card
                    :loading="loading"
                    class="mx-auto my-12"
                    max-width="374"
                  >
                    <v-img
                      height="250"
                      src="https://cdn.vuetifyjs.com/images/cards/cooking.png"
                    ></v-img>

                    <v-card-title>{{ item.Name}}</v-card-title>
                      <v-card-subtitle class="text-left">by {{ item.Teacher }}<br/>{{ item.Year }} {{ item.Semester}} {{ item.DayAt }}</v-card-subtitle>
                    <v-card-text>
                      <v-row
                        align="center"
                        class="mx-0"
                      >
                        <v-rating
                          :value="4.5"
                          color="amber"
                          dense
                          half-increments
                          readonly
                          size="14"
                        ></v-rating>

                        <div class="grey--text ml-4">4.5 (413)</div>
                      </v-row>

<!--                      <div class="my-4 subtitle-1">-->
<!--                        $ • Italian, Cafe-->
<!--                      </div>-->
                        <br/>
                        <div max-height="80px" class="class-description">{{ item.Description }}</div>
                    </v-card-text>

<!--                    <v-divider class="mx-4"></v-divider>-->

<!--                    <v-card-title>Tonight's availability</v-card-title>-->

<!--                    <v-card-text>-->
<!--                      <v-chip-group-->
<!--                        v-model="selection"-->
<!--                        active-class="deep-purple accent-4 white&#45;&#45;text"-->
<!--                        column-->
<!--                      >-->
<!--                        <v-chip>5:30PM</v-chip>-->

<!--                        <v-chip>7:30PM</v-chip>-->

<!--                        <v-chip>8:00PM</v-chip>-->

<!--                        <v-chip>9:00PM</v-chip>-->
<!--                      </v-chip-group>-->
<!--                    </v-card-text>-->

                    <v-card-actions>
                      <v-btn
                        color="deep-purple lighten-2"
                        text
                        @click="reserve"
                      >
                        Reserve
                      </v-btn>
                    </v-card-actions>
                  </v-card>
                </template>

<!--            <template>-->
<!--                <v-card-->
<!--                        class="mx-auto"-->
<!--                        max-width="344"-->
<!--                >-->
<!--                    <v-layout row>-->
<!--                        <v-card-text>-->
<!--                            <p class="display-1 text&#45;&#45;primary">-->
<!--                                {{ item.Name }}-->
<!--                            </p>-->
<!--&lt;!&ndash;                            <p>Campus in {{ faculty.Campus }}</p>&ndash;&gt;-->
<!--                        </v-card-text>-->
<!--                    </v-layout>-->
<!--                    <v-card-actions>-->
<!--                        <v-btn-->
<!--                                text-->
<!--                                color="deep-purple accent-4"-->
<!--                        >-->
<!--                            Learn More-->
<!--                        </v-btn>-->
<!--                    </v-card-actions>-->
<!--                </v-card>-->
<!--                <br/><br/>-->
<!--            </template>-->
        </div>

        <div class="text-center">
            <v-pagination
              v-model="page"
              :length="6"
            ></v-pagination>
        </div>
        <br/><br/><br/>
    </div>
</template>

<script>
    // import SideBar from './SideBar'
    import axios from 'axios'

    var baseURL = "http://localhost:8888"

    export default {
        name: 'ClassList',
        props: {
            msg: String
        },
        data(){
            return{
                items:[],
                loading:false,
                selection:1,
            }
        },
        created() {
            axios.get(baseURL + "/classes")
                .then(response=>{
                    this.items = response.data
                })
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .wrapper{
    }
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
        width:80%;
        height:100%;
    }
    .card{
        /*display:inline-block;*/
        width:50%;
        height:100px;
        margin:5px 7px;
        background-color:green;
    }
    div.class-description {
        display: -webkit-box;
        overflow: hidden;
        text-overflow: ellipsis;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 3;
    }
</style>
