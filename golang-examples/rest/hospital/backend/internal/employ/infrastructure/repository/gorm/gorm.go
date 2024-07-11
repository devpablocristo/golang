
var db *gorm.DB

// CRUD Tabla Employs

// Create (CRUD)
func (e *Employ) Create() {
	db.Exec("ALTER TABLE `personas` AUTO_INCREMENT = 1;")
	db.Save(&e.Person)
	db.Exec("ALTER TABLE `Employs` AUTO_INCREMENT = 1;")
	e.IDPerson = e.Person.IDPerson
	db.Save(&e)
}

// Read (CRUD)
func (e *Employ) Read() {
	var p Person
	db.First(&e, e.IDEmploy)
	db.First(&p, e.IDPerson)
	e.Person = p

	db.Find(&e)

	//fmt.Println("Leer:%d", e.IDEmploy)

	/*
		out, err := json.Marshal(e)
		if err != nil {
			panic(err)
		}

		fmt.Println("\n\n\n\n\nEmploy: " + string(out) + "\n\n\n\n\n")
	*/

	//return e
}

// Read (CRUD)
func (e *Employ) ReadAll() []Employ {
	var todosE []Employ
	var todosPE []Person

	db.Exec("SELECT * FROM Employs e LEFT JOIN personas p ON e.id_persona = p.id_persona ORDER BY e.id_Employ ASC").Find(&todosE)

	db.Exec("SELECT * FROM Employs e LEFT JOIN personas p ON e.id_persona = p.id_persona ORDER BY e.id_Employ ASC").Find(&todosPE)

	for i := 0; i < len(todosE); i++ {
		todosE[i].Person = todosPE[i]
	}

	//c.JSON(http.StatusOK, todosE)
	//c.JSON(http.StatusOK, todosPE)

	return todosE
}

func (e *Employ) Update() {
	var datetime = time.Now()

	datetime.Format(time.RFC3339)

	e.Leer()

	//fmt.Println(e)

	db.Exec("UPDATE Employs SET actualizado_en=? WHERE Employs.id_Employ=?", datetime, e.IDEmploy).Find(&e)
}

func (e *Employ) SoftDelete() {
	var datetime = time.Now()

	datetime.Format(time.RFC3339)

	e.Leer()

	// Enable Logger, show detailed log
	//db.LogMode(true)

	// Debug a single operation, show detailed log for this operation
	//db.Debug().Where("name = ?", "jinzhu").First(&User{})

	//db.Debug().Exec("UPDATE Employs SET eliminado_en=? WHERE Employs.id_Employ=?", datetime, e.IDEmploy).Find(&e)

	db.Exec("UPDATE Employs SET eliminado_en=? WHERE Employs.id_Employ=?", datetime, e.IDEmploy).Find(&e)

	// Disable Logger, don't show any log even errors
	//db.LogMode(false)

}
