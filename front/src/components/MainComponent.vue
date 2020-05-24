<template>
<div class="container">
    <div class="row">
        <div class="col col-md-6 offset-3"> 
            <h1>
                InfoDomain
            </h1>
        </div>
    </div>
    <div class="row">
        <div class="col col-md-6 offset-3"> 
            <div class="input-group mb-3">
                <input type="text" class="form-control form-control-lg" placeholder="domain.com" aria-label="domain.com" aria-describedby="basic-addon2" v-model="domain">
                <div class="input-group-append">
                    <button class="btn btn-outline-secondary" type="button">  <font-awesome-icon  icon="search" v-on:click="getInfoDomain"  v-on:keyup.enter="getInfoDomain"/></button>
                    <button class="btn btn-outline-secondary" type="button"><font-awesome-icon  icon="list-alt"  v-on:click="getRecents"/></button>
                </div>
            </div>
        </div>
    </div>
    <div v-if="servers.length > 0">
        {{servers}} 
    </div>
</div>

   
</template>

<script>
import axios from 'axios'
export default {
    name: 'MainComponent',
    props: {
        msg: String
    },
    methods: {
        getInfoDomain(){
            var self = this
            axios.get('http://localhost:3000/info/'+this.domain)
            .then(function (response) {
                if (response.data != undefined){
                    self.servers = [response.data]
                    //this.servers.push(response.data)
                }
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            });
        },
        getRecents(){
            var self = this
            axios.get('http://localhost:3000/recents')
            .then(function (response) {
            
                if(response.data != undefined){
                    console.log(response);
                    self.servers = response.data
                }
            
            })
            .catch(function (error) {
                console.log(error);
            });
        }
    },
    data(){
        return {
        domain: '',
        servers: []
        
    }
  }
}
</script>