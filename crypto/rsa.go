package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "log"
    "encoding/hex"
    "crypto"
    "bytes"
)

func main()  {
    TestCrypt()
    TestSign()
}

func TestSign() {
    //create private key
    privateKey, err := rsa.GenerateKey(rand.Reader,2048)
    if err != nil {
        log.Fatal("generateKey err ",err)
        return
    }
    // sign
    hashed := sha256.Sum256([]byte("hahaha"))
    /*
    opts := rsa.PSSOptions{
       SaltLength:10,
        Hash:crypto.SHA256,
    }
    */
    sig, err := rsa.SignPSS(rand.Reader,privateKey, crypto.SHA256,  hashed[:], nil)
    if err != nil {
        log.Fatal("SignPSS err ", err)
        return
    }
    log.Println(hex.EncodeToString(sig))
    err = rsa.VerifyPSS(&privateKey.PublicKey, crypto.SHA256, hashed[:], sig, nil)
    if err != nil {
        log.Fatal("VerifyPSS err", err)
        return
    }
    log.Println("VerifyPSS suc")
}

func TestCrypt() {
    //create private key
    privateKey, err := rsa.GenerateKey(rand.Reader,2048)
    if err != nil {
        log.Fatal("generateKey err ",err)
        return
    }
    // encrypt
    ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &privateKey.PublicKey, []byte("hahaha"), []byte("label"))
    if err != nil {
        log.Fatal("encrypt err ", err)
        return
    }
    log.Printf("Ciphertext: %x\n", ciphertext)
    log.Println(hex.EncodeToString(ciphertext))
    plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, []byte("label"))
    if err != nil || bytes.Compare([]byte("hahaha"), plaintext) != 0 {
        log.Fatal("decrypt err: ", err)
        return
    }
    log.Println("plaintext :", string(plaintext))
}

/*
func DecryptOAEPExample()  {
    ciphertext, _ := hex.DecodeString("4d1ee10e8f286390258c51a5e80802844c3e6358ad6690b7285218a7c7ed7fc3a4c7b950fbd04d4b0239cc060dcc7065ca6f84c1756deb71ca5685cadbb82be025e16449b905c568a19c088a1abfad54bf7ecc67a7df39943ec511091a34c0f2348d04e058fcff4d55644de3cd1d580791d4524b92f3e91695582e6e340a1c50b6c6d78e80b4e42c5b4d45e479b492de42bbd39cc642ebb80226bb5200020d501b24a37bcc2ec7f34e596b4fd6b063de4858dbf5a4e3dd18e262eda0ec2d19dbd8e890d672b63d368768360b20c0b6b8592a438fa275e5fa7f60bef0dd39673fd3989cc54d2cb80c08fcd19dacbc265ee1c6014616b0e04ea0328c2a04e73460")
    label := []byte("orders")

    rng := rand.Reader
    key := rsa.PrivateKey{}
    plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, )
}
*/
