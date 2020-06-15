<template>
    <div>
        <div v-if="loading" id="carga-padre">  <!--IMAGE LOADING -->

            <div id="carga"></div>
        </div>
        <div class="container" :class ="loading ? 'load' : 'not-load'" ref="content">
    
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
                    <b-alert :show="error"
                        dismissible
                        fade
                        variant="danger"
                    >
                        Se ha presentado un problema en la consulta
                    </b-alert>
                </div>
            </div>
            <div class="row">
                <div class="col col-md-10 offset-1"> 
                    <div class="input-group mb-3">
                        <input type="text" class="form-control form-control-lg" 
                        placeholder="domain.com" aria-label="domain.com" aria-describedby="basic-addon2"
                        v-on:keydown.enter.prevent="getInfoDomain" v-model="domain"  required>
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

            <br>
            <domain-component
                v-for="domain in servers"
                v-bind:domain="domain"
                v-bind:key="domain.name"
            >
            </domain-component>
        </div>

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
                self.loading = true
                axios.get('http://localhost:3000/info/'+this.domain)
                .then(function (response) {
                    if (response.data != undefined){
                        self.servers = [response.data]
                        //this.servers.push(response.data)
                    }
                    console.log(response);
                    self.loading = false;
                    self.domain = ""
                })
                .catch(function (error) {
                    console.log(error);
                    self.loading = false;
                    self.error = true;
                    self.domain = ""
                });
            }
          
        },
        getRecents(){
            var self = this;
            self.domain = "";
            self.loading = true;
            axios.get('http://localhost:3000/recents')
            .then(function (response) {
                console.log(response)
                if(response.data != undefined ){
                    console.log(response);
                    self.servers = response.data
                    self.loading = false;
                }else if(response.data == null){
                    self.loading = false;
                }
            
            })
            .catch(function (error) {
                console.log(error);
                self.loading = false;
                self.error = true
            });
        }
    },
    data(){
        return {
        domain: '',
        servers: [],
        loading: false,
        error: false
        
    }
  }
}
</script>
<style scoped>
#carga-padre {
	position: fixed;
	top: 2%;
	left: 47%;
}

#carga {
	margin: 300px auto 0 auto;
	border: 16px solid #f3f3f3;
	border-radius: 50%;
	border-top: 16px solid #3498db;
	width: 120px;
	height: 120px;
	-webkit-animation: spin 2s linear infinite;
	/* Safari */
	animation: spin 2s linear infinite;
}

.load {
    opacity: 0.3;
}
.not-load{
    opacity: 1;
}

@-webkit-keyframes spin {
	0% {
		-webkit-transform: rotate(0deg);
	}
	100% {
		-webkit-transform: rotate(360deg);
	}
}

@keyframes spin {
	0% {
		transform: rotate(0deg);
	}
	100% {
		transform: rotate(360deg);
	}
}


</style>