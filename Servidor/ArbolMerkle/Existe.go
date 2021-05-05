package ArbolMerkle

func (this *Arbol) ExisteTienda(raiz *Nodo, hash string, tipo string, indice string, departamento string, nombretienda string, descripcion string, contacto string, calificacion int, logo string) bool{
	if raiz != nil {
		if raiz.Hash == hash && raiz.Tipo == tipo && raiz.Indice == indice && raiz.Departamento == departamento && raiz.NombreTienda == nombretienda && raiz.DescripcionTienda == descripcion && raiz.ContactoTienda == contacto && raiz.Calificacion == calificacion && raiz.Logo == logo {
			return true
		}
		a := this.ExisteTienda(raiz.Izquierda, hash, tipo, indice, departamento, nombretienda, descripcion, contacto, calificacion, logo)
		if a == true {
			return a
		}
		b := this.ExisteTienda(raiz.Derecha, hash, tipo, indice, departamento, nombretienda, descripcion, contacto, calificacion, logo)
		if b == true {
			return b
		}
	}
	return false
}

func (this *ArbolProductos) ExisteProducto(raiz *NodoProductos, hash string, tipo string, tienda string, depa string, calificacion int, nombreProducto string, codigo int, descripcion string, precio float64, cantidad int, imagen string, almacenamiento string) bool{
	if raiz != nil {
		if raiz.Hash == hash && raiz.Tipo == tipo && raiz.Tienda == tienda && raiz.Departamento == depa && raiz.Calificacion == calificacion && raiz.NombreProducto == nombreProducto && raiz.Codigo == codigo && raiz.Precio == precio && raiz.Cantidad == cantidad && raiz.Imagen == imagen && raiz.Almacenamiento == almacenamiento{
			return true
		}
		a := this.ExisteProducto(raiz.Izquierda, hash, tipo, tienda , depa , calificacion , nombreProducto , codigo , descripcion , precio , cantidad , imagen , almacenamiento )
		if a == true {
			return a
		}
		b := this.ExisteProducto(raiz.Derecha, hash, tipo, tienda , depa , calificacion , nombreProducto , codigo , descripcion , precio , cantidad , imagen , almacenamiento )
		if b == true {
			return b
		}
	}
	return false
}

func (this *ArbolPedidos) ExistePedido(raiz *NodoPedidos, hash string, tipo string, fecha string, tienda string, depa string, calificacion int, cliente int, producto int, cantidad int) bool{
	if raiz != nil {
		if raiz.Hash == hash && raiz.Tipo == tipo && raiz.Fecha == fecha && raiz.Tienda == tienda && raiz.Departamento == depa && raiz.Calificacion == calificacion && raiz.Cliente == cliente && raiz.Producto == producto && raiz.Cantidad == cantidad {
			return true
		}
		a := this.ExistePedido(raiz.Izquierda, hash, tipo, fecha , tienda , depa , calificacion , cliente , producto , cantidad )
		if a == true {
			return a
		}
		b := this.ExistePedido(raiz.Derecha, hash, tipo, fecha , tienda , depa , calificacion , cliente , producto , cantidad)
		if b == true {
			return b
		}
	}
	return false
}

func (this *ArbolUsuarios) ExisteUsuario(raiz *NodoUsuarios, hash string, tipo string, dpi int, nombre string, correo string, password string, cuenta string) bool{
	if raiz != nil {
		if raiz.Hash == hash && raiz.Tipo == tipo && raiz.DPI == dpi && raiz.Nombre == nombre && raiz.Correo == correo && raiz.Password == password && raiz.Cuenta == cuenta{
			return true
		}
		a := this.ExisteUsuario(raiz.Izquierda,  hash , tipo , dpi , nombre , correo , password , cuenta  )
		if a == true {
			return a
		}
		b := this.ExisteUsuario(raiz.Derecha, hash , tipo , dpi , nombre , correo , password , cuenta)
		if b == true {
			return b
		}
	}
	return false
}

func (this *ArbolComentarios) ExisteComentarioT(raiz *NodoComentario, hash string, tipo string, Tienda string, departamento string, calificacion int, respondiendo string, dpi int, fecha string, comentarios string) bool{
	if raiz != nil {
		if raiz.Hash == hash && raiz.Tipo == tipo && raiz.Tienda == Tienda && raiz.Departamento == departamento && raiz.Calificacion == calificacion && raiz.Respondiendo == respondiendo && raiz.Dpi == dpi && raiz.Fecha == fecha && raiz.Comentario == comentarios{
			return true
		}
		a := this.ExisteComentarioT(raiz.Izquierda, hash , tipo , Tienda , departamento , calificacion , respondiendo , dpi , fecha , comentarios )
		if a == true {
			return a
		}
		b := this.ExisteComentarioT(raiz.Derecha, hash , tipo , Tienda , departamento , calificacion , respondiendo , dpi , fecha , comentarios )
		if b == true {
			return b
		}
	}
	return false
}

func (this *ArbolComentariosProducto) ExisteComentarioP(raiz *NodoComentarioProducto, hash string, tipo string, Tienda string, departamento string, calificacion int, codigoP int, respondiendo string, dpi int, fecha string, comentarios string) bool{
	if raiz != nil {
		if raiz.Hash == hash && raiz.Tipo == tipo && raiz.Tienda == Tienda && raiz.Departamento == departamento && raiz.Calificacion == calificacion && raiz.CodigoProducto == codigoP && raiz.Respondiendo == respondiendo && raiz.Dpi == dpi && raiz.Fecha == fecha && raiz.Comentario == comentarios{
			return true
		}
		a := this.ExisteComentarioP(raiz.Izquierda, hash , tipo , Tienda , departamento , calificacion, codigoP , respondiendo , dpi , fecha , comentarios )
		if a == true {
			return a
		}
		b := this.ExisteComentarioP(raiz.Derecha, hash , tipo , Tienda , departamento , calificacion, codigoP , respondiendo , dpi , fecha , comentarios )
		if b == true {
			return b
		}
	}
	return false
}
