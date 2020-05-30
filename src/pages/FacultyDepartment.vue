<template>
    <div class="wrapper" align="center">
        <h1>Faculties</h1>
        <v-divider></v-divider>
        <br/>
        <div v-for="(faculty, key, index) in faculties" :key="index">
            <template>
              <v-card
                class="mx-auto"
                max-width="344"
              >
                  <v-layout row>
                <v-card-text>
                  <p class="display-1 text--primary">
                      {{ faculty.Name }}
                  </p>
                  <p>Campus in {{ faculty.Campus }}</p>
                </v-card-text>
                  </v-layout>
                  <v-card-actions>
                     <v-btn
                       text
                       color="deep-purple accent-4"
                     >
                       Learn More
                     </v-btn>
                  </v-card-actions>
              </v-card>
                <br/><br/>
            </template>
        </div>
    </div>
</template>

<script>
    import axios from 'axios'

    var baseURL = "http://localhost:8888"

    export default {
        name: 'FacultyDepartment',
        data(){
            return{
                faculties: [],
            }
        },
        created(){
            axios.get(baseURL + "/faculties")
                .then(response=>{
                    this.faculties = response.data
                })

        }
    }
</script>
