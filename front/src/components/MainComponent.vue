<template>
<div class="container">
    <form v-on:submit.prevent="getInfoDomain" >
    <div class="row">
        <div class="col col-md-6 offset-3"> 
            <h1>
                InfoDomain
            </h1>
        </div>
    </div>
    <div class="row">
        <div class="col col-md-10 offset-1"> 
            <div class="input-group mb-3">
                <input type="text" class="form-control form-control-lg" 
                placeholder="domain.com" aria-label="domain.com" aria-describedby="basic-addon2"
                v-on:keydown.enter="getInfoDomain" v-model="domain"  required>
            </div>
        </div>
    </div>
    <div class="row">
        <div class="col offset-3 col-md-3">
            <button class="btn btn-outline-secondary" type="submit"><font-awesome-icon  icon="search"/><span>Consultar en InfoDomain</span></button>
        </div>
        <div class="col col-md-3">
            <button class="btn btn-outline-secondary" type="button" v-on:click="getRecents"><font-awesome-icon  icon="list-alt"  /> <span>Ver Búsquedas recientes</span></button>
        </div>
        
    </div>
    </form>
 <!--    <div v-if="servers.length > 0">
        {{servers}} 
    </div> -->
    <br>
    <domain-component
        v-for="domain in servers"
        v-bind:domain="domain"
        v-bind:key="domain.name"
    >
    </domain-component>
</div>

   
</template>

<script>
import axios from 'axios'
import DomainComponent from './DomainComponent.vue'
export default {
    name: 'MainComponent',
    components:{
        DomainComponent
    },
    props: {
        msg: String
    },
    methods: {
        getInfoDomain(){
            if(this.domain != ''){
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
            }
          
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