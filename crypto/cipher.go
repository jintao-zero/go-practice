package main

import (
    "log"
    "crypto/aes"
    "crypto/rand"
    "io"
    "crypto/cipher"
    "fmt"
)

func main()  {
    //TestBlock()
    TestStream()
}

func TestStream() {
    // new aes Block, block len should be 16,24,32
    var key [32]byte
    copy(key[:], "haha")
    fmt.Println(len(key))
    block, err := aes.NewCipher(key[:])
    if err != nil {
        log.Println("new cipher err: ", err)
        return
    }
    plaintext := []byte("this is a stream cipher example.")
    ciphertext := make([]byte, block.BlockSize() + len(plaintext))
    iv := ciphertext[:block.BlockSize()]
    _, err = io.ReadFull(rand.Reader, iv)
    if err != nil {
        log.Println("read iv err: ", err)
        return
    }
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[block.BlockSize():], plaintext)
    log.Printf("ciphertext %x", plaintext)

    decryStream := cipher.NewCFBDecrypter(block, ciphertext[0:block.BlockSize()])
    decrytext := make([]byte, len(ciphertext))
    decryStream.XORKeyStream(decrytext, ciphertext[block.BlockSize():])
    log.Println(string(decrytext))
}

func TestBlock()  {
    key := []byte("example key 1234")
    log.Println("key len: ", len(key))
    plaintext := []byte("exampleplaintextexampleplaintext")
    log.Println("plaintext len: ", len(plaintext))

    // CBC mode works on blocks so plaintexts may need to be padded to the
    // next whole block. For an example of such padding, see
    // https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
    // assume that the plaintext is already of the correct length.
    if len(plaintext)%aes.BlockSize != 0 {
        panic("plaintext is not a multiple of the block size")
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    // The IV needs to be unique, but not secure. Therefore it's common to
    // include it at the beginning of the ciphertext.
    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        panic(err)
    }

    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

    log.Printf("%x\n", ciphertext)
    mode = cipher.NewCBCDecrypter(block, ciphertext[:aes.BlockSize])
    pt := make([]byte, len(ciphertext[aes.BlockSize:]))
    mode.CryptBlocks(pt, ciphertext[aes.BlockSize:])
    log.Println(string(pt))
}
