package structural

import "fmt"

/*
The Proxy Pattern is a structural design pattern that provides a surrogate or placeholder for another object to control access to it. It is used to add an additional layer of functionality to an existing object without changing its interface.
There are several types of proxies, such as virtual proxies, protection proxies, and remote proxies.

By using the Proxy Pattern, you can control access to the real object and add additional functionality, such as lazy loading, without changing the interface of the object.
This makes the code more flexible and easier to maintain.
*/

// Image interface
type Image interface {
	Display()
}

// RealImage is the actual image that gets loaded from disk.
type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	loadFromDisk(filename)
	return &RealImage{filename: filename}
}

func (r *RealImage) Display() {
	fmt.Printf("Displaying %s\n", r.filename)
}

func loadFromDisk(filename string) {
	fmt.Printf("Loading %s from disk...\n", filename)
}

// ProxyImage is a proxy for RealImage.
type ProxyImage struct {
	filename  string
	realImage *RealImage
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename: filename}
}

func (p *ProxyImage) Display() {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.filename)
	}
	p.realImage.Display()
}

func ProxyCaller() {
	image1 := NewProxyImage("image1.jpg")
	image2 := NewProxyImage("image2.jpg")

	// The image is loaded from disk only when Display is called.
	image1.Display()
	image1.Display() // This time it won't load from disk
	image2.Display()
	image2.Display() // This time it won't load from disk
}
