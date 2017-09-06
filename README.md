# GopherTrace
A ray tracer written in the Go language, adapted from the C++ methodologies in the book "Ray tracing in one weekend" by Peter Shirley. Check out the book here: https://www.amazon.com/Ray-Tracing-Weekend-Minibooks-Book-ebook/dp/B01B5AODD8.

**What is Ray Tracing?**

Ray tracing is a technique in computer graphics for generating an image by tracing the path of light through pixels in an image plane and simulating the effects of its encounters with virtual objects. Many animation softwares such as Maya are dependent on this technique.

**Objectives of this project**

- To explore the field of Computer Graphics through a hands-on project
- To grasp the fundamentals of the Go language by emulating C++ code, which I am more familiar with.

# How to run it

cd into the project's root directory and run the following command in your terminal:

`go run main.go`

This command will only work if you have Go installed on your system (https://golang.org/dl/). The file returned is a .ppm file, which can be viewed using ToyViewer for OS X, or some other software on other systems. I may add a configuration that returns a png, but haven't gotten there yet.

# Progression of image

1. Rendering metals
![alt text](https://raw.githubusercontent.com/ashwin9798/Gopher-Trace/master/images/currentImage.png)

2. Fuzz property added to metals
![alt text](https://raw.githubusercontent.com/ashwin9798/Gopher-Trace/master/images/image2.png)

3. Dielectric/glass sphere
![alt text](https://raw.githubusercontent.com/ashwin9798/Gopher-Trace/master/images/image3.png)

4. New Camera Angle
![alt text](https://raw.githubusercontent.com/ashwin9798/Gopher-Trace/master/images/image4.png)

5. Random world
![alt text](https://raw.githubusercontent.com/ashwin9798/Gopher-Trace/master/images/image5.png)

6. Motion blur of spheres
![alt text](https://raw.githubusercontent.com/ashwin9798/Gopher-Trace/master/images/image6.png)
