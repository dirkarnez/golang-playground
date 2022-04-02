package main

import "fmt"

type IAdapter interface {
	Read()
	Write()
	Close()
}

type Foo struct {
	IAdapter
}

func NewFoo(adapter IAdapter) *Foo {
	return &Foo{adapter}
}

type AdapterA struct {
	IAdapter // inherits Write() and Close()
}

func (adapterA *AdapterA) Read() {
	// simulate a read
	fmt.Println("AdapterA::Read")
}

type AdapterB struct {
	IAdapter // inherits Write() and Close()
}

func (adapterB *AdapterB) Read() {
	// simulate a read
	fmt.Println("AdapterB::Read")
}

func main() {
	// class Base {
	// 	public String getContent() {
	// 		return "Base";
	// 	}
	// }

	// class Sub extends Base {
	// 	@Override
	// 	public String getContent() {
	// 		return "Sub";
	// 	}
	// }
	// public class Main {

	// 	 public static void main(String []args){
	// 		 Base b = new Sub();
	// 		System.out.println(b.getContent());
	// 	 }
	// }

	var foo_AdapterA *Foo = NewFoo(&AdapterA{})
	foo_AdapterA.Read()

	var foo_AdapterB *Foo = NewFoo(&AdapterB{})
	foo_AdapterB.Read()
}
