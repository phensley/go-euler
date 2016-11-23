package euler

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

// SPOILER WARNING: this file contains the answers!
//
// They are here to enable auto-validation during refactorings
// of common code, and have been obfuscated using a fixed AES key.

func getAnswerCipher() cipher.AEAD {
	block, err := aes.NewCipher([]byte("HideEulerAnswer!"))
	FatalOnError(err, "aes.NewCipher")
	gcm, err := cipher.NewGCM(block)
	FatalOnError(err, "cipher.NewGCM")
	return gcm
}

// HideAnswer encrypts the answer
func HideAnswer(problem, answer string) string {
	gcm := getAnswerCipher()
	nonce := makeNonce(problem)
	ciphertext := gcm.Seal(nil, nonce, []byte(answer), nil)
	return hex.EncodeToString(ciphertext)
}

// RevealAnswer decrypts the answer
func RevealAnswer(problem, secret string) string {
	ciphertext, err := hex.DecodeString(secret)
	FatalOnError(err, "hex.DecodeString")
	gcm := getAnswerCipher()
	nonce := makeNonce(problem)
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	FatalOnError(err, "gcm.Open")
	return string(plaintext)
}

// Converts the problem number into a 12-byte nonce.
func makeNonce(problem string) []byte {
	nonce := make([]byte, 12)
	i := 0
	length := len(problem)
	for i < 12 {
		nonce[i] = byte(problem[i%length])
		i++
	}
	return nonce
}

var answers = map[string]string{
	"001": "502be1a05933f163f22515ee627be65dc193c8ca3413",
	"002": "fb352cd2e1b19f814a6a5d20d13ffac34709d6e1fbfd72",
	"003": "4082d4badbd5c866cdba82a6a1901dc41f04fc4b",
	"004": "0f678c2850f983e78657e5e024a8cb952a1a73e3b88c",
	"005": "d003b78b62c7c055e785deadfcb8271561efa37ec31cd3a261",
	"006": "2066d95312891f9922028b7539b5e6890410d21a08f56a66",
	"007": "56ded3541d4100bc6b2810315e952f27bbe47be68a69",
	"008": "600a9d473c96cc295fb84117492581fcd706d5158691feea0c72e8",
	"009": "f7f813dcd5886ddea93711b6dcba557426f4bba5ae820758",
	"010": "00dbc013bc75948431c8a342ef0cb2a6f2e8405b4b93e76af741203f",

	"011": "de524743767a100382021d955a39f04e87c358eb3a90cab8",
	"012": "5c583d933cff70d14c2ef5068144112c45c35d8e80b2f7c1",
	"013": "28cf93f8422f7c616227b2b3676178bf11d4967ab048e665e2fe",
	"014": "a82102fb9b9d80196e56f36cb61f12aab995b5320b4a",
	"015": "32979cdc7536d24e0882e15fd63d77d8008e28dbd15a588371334558",
	"016": "bb47811cb0169014c0c2bcc686635c46d3e8e270",
	"017": "ae3d9c0879305a7f4eeeb69503cbb6253706b257df",
	"018": "eb1347c6e85388d7cbbadc45a1d81e9f1453fa56",
	"019": "36dec9d84dc71e71bb56d81da5acb3e5396b13",
	"020": "2d1174ad84df4594527044d70d33a260010d82",

	"021": "307c2f0fc9a520ce86e7c2f7519d3ee80562210d31",
	"022": "13fd4c5df900a940736633d59bf8c78c5df545c92f7c30e1fd",
	"023": "a6e99758c87ae571a6f02ca1d6e37a78960219313cfdc0",
	"024": "198780fc1a3f44822671de4f1724b87b26cd77a692df144cf535",
	"025": "ed616e325bd74d6574a7386525f233d8ecc6583f",
	"026": "e95e46c4fc4513dbca3249ace72fa15eccbba4",
	"027": "eea1f2d6cdc6fb5e72e0256d4a2506a942461703456b",
	"028": "e7db14e9ab5599b8d26d9c50b55782d89da529ff9ca77414c0",
	"029": "b8bca2b92694526ada3d266e946a2a923522d3d7",
	"030": "fd5b69aba3fcaf5e9ba3a4447f5168106bd1707cae78",

	"031": "c72278c58dda9b102893e776b236ec42ae3c5658e7",
	"032": "9f01ab99b73b624c2c6f49e1d652e3310712da5b44",
	"033": "80b6bc27efc21dab8abcc67336b95d0e26d5a0",
	"034": "bf49cbcaadcbce3db139cf7050adbd0ce98c3731b7",
	"035": "b4b9c0e4f593753c8e56e2f394d80b44a15b",
	"036": "bfef392d7e38483923a4d5d71ed341df7e6f3a7a805d",
	"037": "17fb4cde4cac46e8f66586628ae8fe98bc91952b76ee",
	"038": "8c9f65548bb471ec8e897dd6f536cd7181bf905d6bb619e8b1",
	"039": "78d74b889468cb3d1e1eaef55b72f625ca9ca7",
	"040": "c4d9e219ae175851506bbf9e65cb533dc79b4a",

	"041": "17ec4eecdc30c81b68bd2c61b2e78a6da8e6cb8d6bec9f",
	"042": "274d0b85f6bbaa36b3d3e309f13d3822e5e1a0",
	"043": "a2e66b6460b81afb892a8747ce2fab87c8b2016fa8638792dc3c54",
	"044": "726329291d5dc4fbd0bcbf839288380649beef9976278a",
	"045": "b27e9d909995c9fb360df08bb4031e82b7f8cbfc26e3d9067c04",
	"046": "83bb7b5ce1a763834cd7762c8f8ab16dde26880d",
	"047": "f26b9c5e03d93f618cd7c5add7b0bff8f3c93b3993be",
	"048": "8af8d34bd7a6c5386a6f998ce4a7d9a44fbb91dcdc04c5fe6c65",
	"049": "0cb1fa2cd687aed2ac8af204e056966ac9af72839cb0430f824c694f",
	"050": "9b3653ac531a4bbe29773204fd6d2a132786306f6139",

	"051": "59835a755e6583c7af2c457056f4080234b20adede0a",
	"052": "d2d83fee41f9443b4d85af7fdf8364d165e833b3c12b",
	"053": "a35aae3b1961a2a161d2fd791d1c10430a109a24",
	"054": "3be2864b0b95f9487a1a7f3be93b0e04c972d2",
	"055": "ccbc78cd7ca0deccf7398b6b09528fbca68637",
	"056": "c2a0db59bedda916b8dfb4648320f4cf893c37",
	"057": "96e2e0b7542131ad127ecd6f8c78c7427bd190",
	"058": "61b5f6f83dcaf7d5006cd9e899403df40d52f516f0",
	"059": "ac1d23ca6eedb453e89b8187f0bef290c902db6ecbc1",
	"060": "910aac704ca3543ac6aa823615de695fb7a6a36943",

	"061": "8d6d2f003d57babf176e894bd0a20e5d7bd2c5bec0",
	"062": "612afffc60253db6248dad547f6a3d61fd0e7f2cec8da77a6b051a87",
	"063": "0801138e857563f2015a27d4b73e6120f176",
	"064": "d761bb79f04651b606dcf94fa752a8246caca6ea",
	"065": "53219fbd98b0b38790289f703d59ff9c5cc1d3",
	"066": "a031a8a21fc7c365d12e9389a7cbee210447ea",
	"067": "3e20a7da67a94e1e5c7efdd7b7cc6f7dfa568973",
	"068": "2ee168072928cf64037e50c98f3fab9d4587eee27d16f8736f2d26966c21b9a1",
	"069": "212a80396ef6908ee3eeb67c1c32ab65771a0b15c3d5",
	"070": "66c953d68f193be505852fd2760597f8814a2fc7105491",

	"071": "cca4c9ae8eddae70b82c50e4d16bf5bb06761bc4870e",
	"072": "319336f714eb8e30393fad38d62fa016d1b46e3d8889c5bffa07bcae",
	"073": "bd6f64f175bf3ed78829fc24fb0a8c298d1c7856b58911",
	"074": "7f7892e67e3d2e308a33daf2ffa1ecd855f2d6",
	"075": "d937d20ed644e66548cbe628aeb22bb910490327b56a",
	"076": "a584cf24fc50b2bc7fef6eef0432252a48063c755295cdf352",
	"077": "014f9df9d6d5a1ea82fae363109441e1e180",
	"078": "a0183c9e2f6687273fce4e0e0c146a7ccd55123b32",
	"079": "c4ec401f7bcd30cc77f9b1514e1ff64b5666726a74243ff2",
	"080": "c52b8685d5be447000a63bcaf6547516bd5b9a2e5c",

	"081": "b8ae7ab9208bf1426509f09d937c1faf15690ca0c0df",
	"082": "f99360621bfe764e72e6d7cd76203acbd524ab75107e",
	"083": "74c451d9beb804641ee04d0718b3b4827112b976004c",
	"084": "583959916329fb1a8bd1e7f3735e8b955c56ef402e76",
	"085": "f74e155328a17ec7d5a19810762b8d96bab908b7",
	"086": "f27d3abb465db3591d3e1a69319bf0482209f699",
	"087": "934f23ee033c702b355e620f3e6f94d21693b4067473dc",
	"088": "1dac627fce9fcea08a8bda215c6fcce44abb7b8518e1e3",
	"089": "4f892da0e7abbfd9e48d26492c0a083dd9924d",
	"090": "8657c1f0ddd9e282408e5fedf988df297d115858",

	"091": "ae4d4bd96034c9c85779a68469615ecc9924f95fe8",
	"092": "4b23f6c06cf8f3cfa10497bbd6a32e0ef035108737e830",
	"093": "8bff151d2d8e2e97bd566fc243447cb044a6695b",
	"094": "485cbaae4a3e8cb159c79d6d98b53b37eb35de801d1ef7c7e1",
	"095": "11357b6cc91922755ed21e55108b229a439e2810f2",
	"096": "e3d9f266fb93d03f4239db596488f2d03c92dc6488",
	"097": "3838c0a0a6662f8cbbccc54821f4b9ab8122196c79540de638dd",
	"098": "a6c72b721b276bf30b7dd0185ef4f831721051baeb",
	"099": "5fc4d1b75e149dd3b97e5dfbf80dff7ecd68e9",
	"100": "655638e9d7b659c3255af2b1c6007b78f8a3b1f76785c95ce3ea2b1a",
}
