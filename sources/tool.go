/**
* This file is licensed under BSD 3-Clause License
*
* Copyright (c) 2020, magnar-blj
* All rights reserved.
*
* Redistribution and use in source and binary forms, with or without
* modification, are permitted provided that the following conditions are met:
*
* 1. Redistributions of source code must retain the above copyright notice, this
*    list of conditions and the following disclaimer.
*
* 2. Redistributions in binary form must reproduce the above copyright notice,
*    this list of conditions and the following disclaimer in the documentation
*    and/or other materials provided with the distribution.
*
* 3. Neither the name of the copyright holder nor the names of its
*    contributors may be used to endorse or promote products derived from
*    this software without specific prior written permission.
*
* THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
* AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
* IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
* DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
* FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
* DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
* SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
* CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
* OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
* OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*
*/

package main // define current package as main

import ( //import some utils
	"fmt" // import the fmt librairy as stdio as in c 
    "os" // import os module for os type scanning
    "path/filepath" // import filepath librairy for file location scanning
    "math/rand" // import random mathlib to generate random stuff xD
    "log" // import log for logging utils
)

func Walk(absoultePath string, do func(path string)) { // begin function walk
    var files []string // define files as string
    var err error // define err as error type
    
    if !filepath.IsAbs(absoultePath) { // check if argumented absolute file path is provided
        absoultePath, err = filepath.Abs(absoultePath) // refer from err to absFilepath
        if err != nil{ // if err isnt nil do...
            log.Fatal(err) // fatal error message
            return  // return null as default
        }
    }
    
    fmt.Println("walking ...", absoultePath) // some status progress message
    if Exists(absoultePath) == false{ // check if path exist
        log.Fatal(absoultePath, "is not exists directory.") // error message
    }

    err = filepath.Walk(absoultePath, func(path string, info os.FileInfo, err error) error { // create object err for os file infor and stuff xD
		if !info.IsDir(){ // check dir allocation
			files = append(files, path) // append dir to files
		} 
        return nil // return null as global 
    }) /// end
     
    if err != nil { // if err isnt nil do...
        log.Fatal(err) // fatal error message
        return // return null as default
    }
    for _, path := range files { // for path in files do...
		do(path) // check existence
    }
}

func Exists(name string) bool { // begin function exists with param name
    if _, err := os.Stat(name); err != nil { // if name and os.Ststname doesnt match do...
        if os.IsNotExist(err) { // if os doesnt exist do...
            return false // return existence to false
        }
    }
    return true // esle return existence to true
}

func StringInSlice(a string, list []string) bool { // begin StringInSlice with param a as string and list as typecasted array and refer the actual function to bool type
    for _, b := range list { // for name in b of range in list do....
        if b == a { // if b equals to a do
            return true // return true as global
        }
    }
    return false // esle return false as global
}
var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") // all exisitng letters in a file

func randSeq(n int) string {  // begin function randSeq with param n as int and refer the actual function as string type
    b := make([]rune, n) // declare b as rune value
    for i := range b { // for i in rango of b do....
        b[i] = letters[rand.Intn(len(letters))] // define b with param i recusrively and generate random letters with scanned length
    }
    return string(b) // return  typecasted string from b
}

func getdrives() (r []string){ // begin function getDrives 
    for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ"{ // possible drive names
        _, err := os.Open(string(drive)+":\\") // define name of err to unexisting driver
        if err == nil { // if err isnt nill do...
            r = append(r, string(drive)) // append string to drive
        }
    }
    return // return null as default
}

