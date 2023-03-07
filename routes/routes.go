package routes

import (
	//"database/sql"
	"fmt"
	"go_module/database"
	"go_module/render"
	//"log"
	"net/http"

	//_ "github.com/jackc/pgx/v5/stdlib"
)



func Home(w http.ResponseWriter , r *http.Request){
	render.RenderTemplate(w,"home.page.tmpl")
}
func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,"about.page.tmpl")
}
func General(w http.ResponseWriter , r *http.Request){
	render.RenderTemplate(w,"general.page.tmpl")
}
func Major(w http.ResponseWriter , r *http.Request){
	render.RenderTemplate(w,"major.page.tmpl")
}
func Contact(w http.ResponseWriter , r *http.Request){
	render.RenderTemplate(w,"contact.page.tmpl")
}
func Data(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,"data.page.tmpl")

}
func Booking(w http.ResponseWriter, r *http.Request){
	first:=r.FormValue("first_name")
	last:=r.FormValue("last_name")
	email:=r.FormValue("email")
	start:=r.FormValue("start_date")
	end:=r.FormValue("end_date")
	phone:=r.FormValue("phone")

	fmt.Fprintf(w,"First Name is %s",first+"\n")
	fmt.Fprintf(w,"Last Name is %s",last+"\n")
	fmt.Fprintf(w,"Email is %s",email+"\n")
	fmt.Fprintf(w,"Starting date is %s",start+"\n")
	fmt.Fprintf(w,"End Date is %s",end+"\n")
	fmt.Fprintf(w,"Phone Number is %s",phone+"\n")


	/*conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=Bookings_demo user=postgres password=password")
	if err != nil {
		log.Fatalf(fmt.Sprintf("Unable to connect: %v\n", err))
	}

	defer conn.Close()

	log.Println("Connected to database")

	err=conn.Ping()
	if err!=nil{
		log.Fatal("Cannot Ping the database")
	}
	log.Println("pinged database")

	query:=`insert into users(first_name,last_name,phone_number,email,arrival,departure) values($1,$2,$3,$4,$5,$6)`

	_,err=conn.Exec(query,first,last,phone,email,start,end)

	if err!=nil{
		log.Fatal(err)
	}

	log.Println("Inserted a row")*/

	database.InsertRow(first,last,email,phone,start,end)


}
func PostQuery(w http.ResponseWriter, r*http.Request){
	

}
