package main

import (
	"bytes"
	"time"
	"encoding/binary"
	"log"
	"crypto/sha256"
)

//区块结构体
//定义结构
type Block struct {
	//版本号
	Version uint64
	//前区块哈希
	PrevHash []byte
	//梅克尔根
	MerKleRoot []byte
	//时间戳
	TimeStamp uint64
	//难度值
	Difficulty uint64
	//随机数
	Nonce uint64


	//当前区块哈希
	Hash []byte
	//数据
	Data []byte

	//添加区块
	//重构代码
}

//创建区块
func NewBlock(data string,prevHash []byte)*Block  {
	block := Block{
		Version:00,
		PrevHash:prevHash,
		MerKleRoot:[]byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty:100,
		Nonce:100,
		Data:[]byte(data),
	}
	block.SetHash()//设置hash值
	return &block
}

//创建创世块
func NewGenesisBlock(data string,prevHash []byte) *Block {
	block := Block{
		Version:00,
		PrevHash:prevHash,
		MerKleRoot:[]byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty:100,
		Nonce:100,
		Data:[]byte(data),
	}
	block.SetHash()//设置hash值
	return &block
}

//生成哈希
func (b *Block)SetHash()  {
	temp := [][]byte{
		Unit64ToByte(b.Version),
		b.PrevHash,
		b.MerKleRoot,
		Unit64ToByte(b.TimeStamp),
		Unit64ToByte(b.Difficulty),
		Unit64ToByte(b.Nonce),
		b.Data,
	}
	data := bytes.Join(temp,[]byte{})
	hash := sha256.Sum256(data)
	b.Hash = hash[:]
}

//uint64转[]byte函数
func Unit64ToByte(data uint64)[]byte  {
	var buffer bytes.Buffer
	//将数据以二进制的形式存储到buffer中
	err := binary.Write(&buffer,binary.BigEndian,data)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}