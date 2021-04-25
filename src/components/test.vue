
<template>
<v-app>
  <v-data-table
    :headers="Headers"
    :items="doctors"
    :items-per-page="10"
    align="center"
  ></v-data-table>
  <v-row justify="center">
    <v-btn
      color="primary"
      dark
      @click.stop="adddialog = true"
    >
      add Doctor
    </v-btn>


    <v-dialog
      v-model="adddialog"
      id="adddialog"
      name="adddialog"
           
    >
      <v-card>
        <v-card-title class="headline">
          Add doctor
        </v-card-title>

        <v-card-text>
          fill the form to add a doctor
        </v-card-text>
      <form class="form" v-on:submit.prevent="createDoctor">
        
        <v-text-field

                  id="name"
                  label="name"
                  v-model="form.name"
                  required
                >
        </v-text-field>
        <v-text-field
                  id="years_of_experience"
                  label="years_of_experience"
                  v-model="form.years_of_experience"
                  required
                >
        </v-text-field>
         <v-text-field
                  id="type_of_doctor"                  
                  label="type_of_doctor"
                  v-model="form.type_of_doctor"
                  required
                >
        </v-text-field>
         <v-text-field
                  id="consultation_fee"
                  label="consultation_fee"
                  v-model="form.consultation_fee"
                  required
                >
        </v-text-field>
         <v-text-field
                  id="mail_id"
                  label="mail_id"
                  v-model="form.mail_id"
                  required

                >
        </v-text-field>
         <v-text-field
                  id="phone"
                  label="phone"
                  v-model="form.phone"
                  required
                >
        </v-text-field>
        
        <v-card-actions>
          <v-spacer></v-spacer>

          

          <v-btn
            color="green darken-1"
            text
            
            @click="createDoctor()"
          >
            submit
          </v-btn>
          <v-btn
          @click="adddialog = false;
         
            ">
            close
          </v-btn>

          
        </v-card-actions>
        <v-btn
        
        

        >
        </v-btn>
      </form>
      </v-card>
    </v-dialog>


  </v-row>
  <template v-slot:actions="{ item }">
      <v-icon
        small
        class="mr-2"
        @click="editItem(item)"
      >
        mdi-pencil
      </v-icon>
      <v-icon
        small
        @click="deleteItem(item)"
      >
        mdi-delete
      </v-icon>
    </template>=
</v-app>
</template>
<script>
import doctorServices from "@/services/doctorService" 
import axios from 'axios'
axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';


  export default {
    data :() => ({

        form :{
          name:'',
          years_of_experience:'',
          type_of_doctor:'', 
          consultation_fee:'',
          mail_id:'',
          phone:'',
          hospital_id:'55555',


        }
      ,
      updateDoctor :{

      },
      adddialog: false,
 

        Headers:[
            {text:'doctor id', value:"id"},
            {text:'doctor name', value:"name"},
            {text:'type of doctor', value:"type_of_doctor"},
            {text:'years of experience', value:"years_of_experience"},
            {text:'consultation fee', value:"consultation_fee"},
            {text:'email_id', value:"mail_id"},
            {text:'phone numbers', value:"phone"},
            { text: 'Actions', value: 'actions', sortable: false },
            

            
        ],
        doctors:[] ,
        
        

    }),
    created(){
        console.log('inside created')
        this.initialize()
    },
    methods:{
        initialize()
        {this.loadDoctors()},
        
        



        loadDoctors(){

            this.doctors=[]
            doctorServices.getAlldoctors()
            doctorServices.getAlldoctors().then(item=>{
                    item.doctor.forEach(
                        doctor=>{this.doctors.push(doctor)}
                    )
                    
                

            });
            

        },
        createDoctor(){
          axios.post('http://127.0.0.1:5000/doctors', this.form)
          .then((res)=>{
            console.log(res);
            if (res.status==200){
              this.adddialog = false
              this.loadDoctors()
            }
          })
          .catch((err)=>{
            console.log(err);
          })
          .finally(()=>{})
        },
        
        
          
        }
    }
  
</script>

  