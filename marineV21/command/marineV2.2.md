## MarineVersion 2.2

### CONFIGURE

+ Create Channel

        ./1.createCh peerorg0 orderer0.ordererorg0 starpoly Admin ../../localtest/channel-artifacts/starpoly.tx    
    
+ Update Channel (Anchor)

        ./2.updateCh peerorg0 orderer0.ordererorg0 starpoly Admin ../../localtest/channel-artifacts/peerorg0.ordererorg0.starpoly.anchor.tx   
        ./2.updateCh peerorg1 orderer0.ordererorg0 starpoly Admin ../../localtest/channel-artifacts/peerorg1.ordererorg0.starpoly.anchor.tx   
    
+ Join Channel
        
        ./3.joinCh peerorg0 orderer0.ordererorg0 starpoly Admin  
        ./3.joinCh peerorg1 orderer0.ordererorg0 starpoly Admin 
    
+ Install chaincode
    
        ./installChaincode peerorg0 starpoly Admin cat v0 github.com/utxo  
        ./installChaincode peerorg1 starpoly Admin cat v0 github.com/utxo  
        
        
        ./installChaincode peerorg0 starpoly1 Admin cat v0 github.com/cat2
        ./installChaincode peerorg1 starpoly1 Admin cat v0 github.com/cat2
        
+ Instantiate  
        
        ./5.instantiateCC peerorg0 starpoly Admin cat v0 github.com/utxo '{"Args":["init","starpoly"]}' 1  
        
        
        ./5.instantiateCC peerorg0 starpoly Admin utxo v0 github.com/utxo '{"Args":["init","starpoly"]}' 1  
        
        
+ Upgrade (Install -> instantiate -> sync by using invoke-update)

        CAT
        ========================================================================
        ./installChaincode peerorg0 starpoly Admin cat v15 github.com/cat2
        ./installChaincode peerorg1 starpoly Admin cat v15 github.com/cat2
        ./instantiateCC peerorg0 starpoly Admin cat v15 github.com/cat2 '{"Args":["init","starpoly"]}' 1
                
        [invoke]
        ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["create","paul003",{"owner":"paul"}]}' att
        ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["create","paul002",{"owner":"paul"}]}'
        
        [update]
        ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["update","paul001",{"owner":"paul-noPad"}]}'
        ========================================================================
        
        PRIVATE
        ========================================================================
        ./installChaincode peerorg0 starpoly Admin private v4 github.com/cat2
        ./installChaincode peerorg1 starpoly Admin private v4 github.com/cat2
        ./instantiateCC peerorg0 starpoly Admin private v4 github.com/cat2 '{"Args":["init","starpoly"]}' 1
        
        
        [invoke]
        ./marineV2 8 localhost:50061 Admin starpoly private '{"Args":["create","paul001",{"owner":"paul"}]}' att
        [update]
        ./marineV2 8 localhost:50061 Admin starpoly private '{"Args":["update","paul001",{"owner":"paul"}]}' att
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","paul001"]}'
        ========================================================================
        
        DID
        ========================================================================
        ./installChaincode peerorg0 starpoly Admin did v5 github.com/did
        ./installChaincode peerorg1 starpoly Admin did v5 github.com/did
        ./instantiateCC peerorg0 starpoly Admin did v5 github.com/did '{"Args":["init","starpoly"]}' 1
        ========================================================================
                
        
        UTXO
        ========================================================================
        ./installChaincode peerorg0 starpoly Admin utxo v1 github.com/cat2
        ./installChaincode peerorg1 starpoly Admin utxo v1 github.com/cat2
        ./instantiateCC peerorg0 starpoly Admin utxo v1 github.com/cat2 '{"Args":["init","starpoly"]}' 1
        ========================================================================

        UTXO - STARPOLY1
        ========================================================================
        ./installChaincode peerorg0 starpoly1 Admin utxo poly1v1 github.com/cat2
        ./installChaincode peerorg1 starpoly1 Admin utxo poly1v1 github.com/cat2
        ./instantiateCC peerorg0 starpoly1 Admin utxo poly1v1 github.com/cat2 '{"Args":["init","starpoly1"]}' 1
        ========================================================================
        
        ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["update","key",{"owner":"abc2"}]}' empty
        ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["update","paul001",{"owner":"paul"}]}' att
        
        ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["create","paul001",{"owner":"paul"}]}' att

        ./marineV2 8 localhost:50061 admin starpoly cat '{"Args":["cc2cc", "", "cat", "query", [{"field":"address","operator":"=","operand":"z4MTiNvHs76DpwSH3gpivBuvuDynuhefy"}]]}'
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector":{"owner":{"$gt":null}}}]}'
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","key"]}'
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","key1"]}'
    
    
## BASE1 FOR Admin

        
+ INVOKE
        
        ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["create","key1",{"owner":"abc"}]}' empty
        ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["update","key",{"owner":"abc1"}]}' empty
        
        ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["create","A0001", {"doctype":"TA", "docDate":"2017-07-28"}]}' pad
        
        ./marineV2 8 localhost:50061 Admin starpoly utxo '{"Args":["create","A0001", {"doctype":"TA", "docDate":"2017-07-28"}]}' pad
        
        ./marineV2 8 localhost:50061 Admin starpoly1 utxo '{"Args":["create","A0001", {"doctype":"TA", "docDate":"2017-07-28"}]}' pad
        
        

+ QUERY

        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","key"]}'
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","paul002"]}'
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"owner","operator":"=","operand":"abc"}]]}'
        
        ./marineV2 7 localhost:50061 Admin starpoly utxo '{"Args":["query",[{"field":"doctype","operator":"=","operand":"TA"}]]}'
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"doctype","operator":"=","operand":"TA"}]]}'

        [Any]
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector":{"owner":{"$gt":null}}}]}'
        
        ./marineV2 7 localhost:50061 Admin starpoly1 cat '{"Args":["queryByPass",{"selector":{"doctype":{"$gt":null}}}]}'
        ./marineV2 7 localhost:50061 Admin starpoly1 utxo '{"Args":["queryByPass",{"selector":{"doctype":{"$gt":null}}}]}'


## BASE2 FOR Admin


+ INVOKE
        ./marineV2 8 localhost:50061 Admin starpoly1 cat '{"Args":["create","TEST001", {"doctype":"TA", "docDate":"2017-07-28"}]}' pad
        ./marineV2 8 localhost:50061 Admin starpoly1 cat '{"Args":["create","TEST002", {"doctype":"TA2", "docDate":"2017-07-28"}]}' pad
         
         ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["create","PJ001>PJ002", {"doctype":"TA", "docDate":"2017-07-28"}]}' pad
    
         ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["create","PJ001.PJ002", {"doctype":"TA", "docDate":"2017-07-28"}]}' pad
    
         ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["create","PJ001@PJ002", {"doctype":"TA", "docDate":"2017-07-28"}]}' pad
         
         ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["create","PJ001<PJ002", {"doctype":"TA", "docDate":"2017-07-28"}]}' pad
        
         ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["create","PJ001(PJ002", {"doctype":"TA", "docDate":"2017-07-28"}]}' pad
    
+ QUERY
    
         ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"docDate","operator":"=","operand":"2017-07-28"}]]}'
         
         ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"docDate","operator":"=","operand":"2017-07-28"}]]}'
        
         ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query","PJ001>PJ002"]}'
        
         ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query","TEST001"]}'
         
         ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","PJ001>PJ002"]}'  --> "PJ001\u003ePJ002"
         ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","PJ001<PJ002"]}'  --> "PJ001\u003cPJ002"
         ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","PJ001\u003cPJ002"]}'  --> "PJ001\u003cPJ002"
        
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryAttachment","paul001"]}'
        
        
        ./marineV2 7 localhost:50061 Admin starpoly1 cat '{"Args":["queryByKey","TEST001"]}'
        ./marineV2 7 localhost:50061 Admin starpoly1 cat '{"Args":["query",[{"field":"doctype","operator":"=","operand":"TA8"}]]}'
        
        ./marineV2 7 localhost:50061 Admin starpoly1 cat '{"Args":["queryByPass",{"selector":{"doctype":{"$gt":null}}}]}'

+ UPDATE

        ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["update","PJ001", {"doctype":"TA2", "docDate":"2019-07-28", "MoDate":"2017-07-28"}]}' ABBCC
    
         ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["update","PJ001", {"doctype":"TA2", "docDate":"2019-07-28", "MoDate":"2017-07-28"}]}' CBA


+ Delete

        ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["delete","btc"]}' opt
        
        ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["delete","key1"]}' opt
        
        ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["delete","c2aaeb17f158f25d73449b8731710b2ce0cd7c7157ea53f44d4975c5000d68ce"]}' opt

+ historyByKey

        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["getHistoryForCIbyKey","z4MTiNvHs76DpwSH3gpivBuvuDynuhefy@btc"]}'
    

## Decentralized ID (DID)

+ create
    
        ./marineV2 8 localhost:50061 Admin starpoly did '{"Args":["createDidDoc","did3","documents"]}'

+ read

        ./marineV2 7 localhost:50061 Admin starpoly did '{"Args":["readDidDoc","did2"]}'

+ update(undergoing)

        ./marineV2 8 localhost:50061 Admin starpoly did '{"Args":["updateDidDoc","did3","newDIDdoc","proofvalue"]}'

+ delete

        ./marineV2 8 localhost:50061 Admin starpoly did '{"Args":["deleteDidDoc","did2"]}'



## PVT FOR Admin

**PREREQUISITES**

+ Chaincode
    + A chaincode named with **[private]** MUST be installed and instantiated in the same channel 

+ Establishment
    + The membershipcheck() is required based on the membership in the putPrivate() &#10137; putPrivate_()


+ putPrivate()

        ./marineV2 111 localhost:50061 Admin starpoly cat '{"Args":["putPrivate","TITLE", "KEY", "VALUE"]}'  (env added)
    
        ./marineV2 111 localhost:50061 Admin starpoly cat '{"Args":["putPrivate","btc", "BTCPVT001", {"privatedata":"paulbtc", "docDate":"2019-07-28"}]}'
        (priv--> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)

        
        CHECK --> (This will be REMOVED!!)
        ./marineV2 7 localhost:50061 Admin starpoly private '{"Args":["queryByKey","BTCPVT001"]}'
        
    **[VERY IMPORTANT NOTICE]**
    
        The chaincode "PRIVATE" includes only functions related "private", which does not allow any normal accesses and even admin queries like type[7].
        IT MUST BE CALLED BY/FROM ITS MOTHER CHAINCODE.
        
+ getPrivate()

        ./marineV2 111 localhost:50061 Admin starpoly cat '{"Args":["getPrivate","TITLE", "KEY"]}'  (env added)

        ./marineV2 111 localhost:50061 Admin starpoly cat '{"Args":["getPrivate","btc", "BTCPVT001"]}'
        (priv--> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","BTCPVT001"]}'

        ./marineV2 111 localhost:50061 Admin starpoly cat '{"Args":["getPrivate","btc", "CATFEE"]}'
        (priv--> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)
        
        
## CAT FEE

**PREREQUISITES**

MUST ESTABLISH FIRST ON [TITLE] that will be used for "CATFEE", actually "CATFEE" is going to be reserved for FEE title

+ putFEE

    ./marineV2 111 localhost:50061 Admin starpoly cat '{"Args":["putFEE",TITLE, KEY, VALUE]}' (added ENV.)

    ./marineV2 111 localhost:50061 Admin starpoly cat '{"Args":["putFEE","btc", "CATFEE", {"transferfee":"0.01", "destaccount":"z2UPj4Vr2pvREww7jjX1SDoqLtix7toVN", "label":"btc"}]}'
        (priv--> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)

    ./marineV2 111 localhost:50061 Admin starpoly cat '{"Args":["putFEE","btc", "CATFEE", {"transferfee":"0.01", "destaccount":"z2UPj4Vr2pvREww7jjX1SDoqLtix7toVN", "label":"btc"}]}'
        (priv--> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)

+ getFEE
    
    EX) ./marineV2 111 localhost:50061 Admin starpoly cat '{"Args":["getPrivate","TITLE", "KEY"]}'  (env added)
    
    ./marineV2 111 localhost:50061 Admin starpoly cat '{"Args":["getFEE","btc", "CATFEE"]}'
        (priv--> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)


+ chargeFee (807)
    
    ./marineV2 807 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW z4MTiNvHs76DpwSH3gpivBuvuDynuhefy 8.86 etc
            (priv -> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)

+ check (807)
    
    ./marineV2 888 localhost:50061 z2UPj4Vr2pvREww7jjX1SDoqLtix7toVN
    
    
## LOGIN

    ./marineV2 108 localhost:50061 Admin starpoly cat


## PKCS7
    
    ./marineV2 301 passwd text
    
    ./marineV2 302 passwd encryptedMsg



## CROSS ASSET TRANSFER (CAT)


#### Key Generation with Address
    ./marine 88 (option:seed)
    
        EDWARD USER 1   (Paul)  
            > Public  Key ---> 6xr1SFx4Vci8FgzEu1wnGqfGSqy3uaTC9byio13xX9R7
            > Private Key ---> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU
            > Address -------> z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW

        EDWARD USER 2   (Jay)
            > Public  Key ---> fTPAAQC2jC9cxJtqKhsbs77qhgoxeFTaw66HQkUdqRq2
            > Private Key ---> 5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP
            > Address -------> z4MTiNvHs76DpwSH3gpivBuvuDynuhefy

        EDWARD USER 3 (Tim)
            1) Public  Key ---> 3CVxYtwZfsR5UiD7S2Sf4NZx2eoA8SiVRRPwzPTcbGr6
            2) Private Key ---> 38UJ99hwPzmLZbxhmcXhuDoWZEbaGF4g6U78TuYdHjYi1pCZehpvVPuCrW7G9zh2cMzQpmK7z5xrPgCaQpkMUiMa
            3) Address -------> zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3


        ECC USER4
            1) Public  Key ---> 1rQ9dSnnVgGxXxYQa9WC2DCUeay5awvFisAsizd4AeP55m7LbMU5xS9VNNHiEUDtfqtC3kRpJcqa9KApnfCaK7nePUSrUgGxNbv3W2v2p7wkPgXqENPTvmCSXTdM
            2) Private Key ---> vDJgRQ3rKt2uZCSmR5QdKt6eGp7zFT2zDTAJTs52AMdjLiuowQtw3ea2wT87Cf8Lwqe2WUAoc8png3mp764f4HNXRdH6n3T5Zo1tf7AAaoBf65Fhek4vVmDGjxUHfc9ksXvgXRahruYYFGxU4Er1pYWnnyU5dZmjwZxwr
            3) Address -------> z2UPj4Vr2pvREww7jjX1SDoqLtix7toVN 


        RSA USER5
            1) Public  Key (base58) ---> 4DzbtsFgbEQuvNmmADDPeEu5tpHJj4381KmMV2GHbF9pkp8DnFbErSAkgzNk5n97U57ictYbMcwSNE5s4DckfL2M4NUPJsHd8b4badwtHnR6vrszPTGCnBpWLQxF7N1jpMe5NQA35W24S8mGXp1RfmtCbtvDNLLPAPscFWsoEBVL12m5cw8kBgKcJtSykxH4zyZCemtntUUSgRg6o81psiggWrYXSEUC7TE8axwWQMi7ZjnYR2PsvEmgW9jmmp2NiePuxWHKnQA9XALLUigpBWSqqpfAKstT9mAfqcvCuJBx5mwm7N1DBt5dYEbwgNnU3EeByBtCSfwTB65vbRzrB7AKLhoU319PvYqafSM1Jt54TjPBp
            2) Private Key (base58)---> 2rm2QNNaHzhtdRY7bCNXK9Wu7P8JUx5WrpUq6yAh7J6Aonukp4NgNEyE3p4w6xyvErPrVoPWA8EryazsGqUbAxEY9k8TuEro2J5MedrBdbq17AapWMKpZwsmes1EDuq1iDSGtjftfpAr6gLFyVJDfzo2WeHpUshxW467D2vHxQyqdTmkkKPEYRjh1TfoZxUU1e7xCinoNEfWfduz11E2BciSN2sX9h9YEhWRQj4nPwRs9AsoHykAA2USM3ru3fv1DCTGea5mVLyr6kTjqBrGSiCARmdBczGQD7P2ngstcyLKmSEA6mh8xF63Xwb5aeZRTmWJrqKc1yGXtM9uwuZjoT3raH7B1BP12ygnGw2rABdnWBZaNWuWxaKoTquMiQeq5jZRDduUGypG4aWNwL5RCWFkJEDAs7gdvTtWZDwxvbY4shjxWLCpQrPVAtG3zs7awWrP92tBGhJqfft7rojahfqhjP8UkkKSqJCCD8j5LyDgLMxRbiScUuorGbWCSSKo9XXZW7qCB5DwyV73v2qpMLgMxrLrFBrEzCdz585rBupJza2pZk4wab7xqEVwmKqj76UyDca6TaE89XSHL2XBBXKiPmTVneorGc746pgC12LXFkw8fFE2Gas6Nk4LQaj662iqSwRh7z1QSN4t6QdgMVxP2JfDfRoBHJJCSwHNUVQyCihjWmzyAQjm12yZLLwFxf8ZbSEzwyZH62Umb1XgQMh8Nzu2hXjMHfnjp852gbCf3uR3MBhGFNSG3rq5L5X46oUDUYZCkCR85ySXB3pPVmDaYUpunvD7QRbc1c5yuZZdb4fbYfbrKLENrMSp6GHYg3UCZk3R9uPFPtugNDZmovf4Fu2JV9WH6Mn8x6paCryXzk4W9gBj1rgvq3CSxqKdPHJUdkNeaKaQBxUd3U9HZrSXvXoH7CPFDZwceSh5FKaqqMtpucXRvDM3VxV6cgJFFTjt8GEzQeBaouexoMZdV52Fkermx6wQpU5RTckFxTo5zEjqLQiWFRtHWi1zUUDxp1XdGAFZELKju1eBKWrdWo3g1G58DCouxGpRkbXmVn2nK1X4jjbHDVk1woQY2EYmJmt8D6kwBhWTbgnVpqDNdDCYsZ4cUv6bPXWGtym4YHXxS9Ekdww3vMkhfZPcjVDtDbeh6UcFLQbSkKFGRXk3io2fg5iRLDWGDGAVtZM6HvSc6MUvf8F3Y3e4Diiuhrp9BZCcTfXGdX4ssR9y3sBxBAnYmdWemiLvrUauxRUnbnR7bHocs2N4EfJkrv794HRp5NjuCL3AGt8RT8jah3EDTTEf2oUxqVhGSP31nfV2RWwmdytp2XvGHvD2r5xGtA3wkG2k8xzebgvnxJC7amUXQfNXSwmidaJd2updYmZHxkdjaday9RJcd9GXoYozfu5c15uoH4U8U2ybqUc83WQdd7JfWzbmR6vDgRdhg9iqSREgugtmEhYmSPqMAAC2E6Jd4Se7buqhaMNdmRvUKoiHNvb6kNTz7MMbMrZhBGG8SFTYSkBewcLTR9NhCFTsddNt3Wn7P2NFS4PeUtix4VvbZEhPE1YzagpoUpzbvUAsB6G7NKqvCX2aYErjTfAH6RMQtdVHe55xoB6t9EFY9wCW6YoApxF
            3) Address -------> z3W7tjpeVMkgr8EjVbCjsUcYNSh6t8fxs



#### balanceOf

    [BTC Issuer]
    ./marineV2 888 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW    
    
    [ETC Issuer]
    ./marineV2 888 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy
    
    ./marineV2 888 localhost:50061 zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3
    
    ./marineV2 888 localhost:50061 z5cXux681QCFJqYdTaDju7Hp4NFQ7jCZo
    
    ./marineV2 888 localhost:50061 zkRuxEW3dEAyGBZaSEbdc8ZqseUUjvxo
    
    (ECC User)
    ./marineV2 888 localhost:50061 z2UPj4Vr2pvREww7jjX1SDoqLtix7toVN
    
    (RSA User)
    ./marineV2 888 localhost:50061 z3W7tjpeVMkgr8EjVbCjsUcYNSh6t8fxs
    
    
    marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"address","operator":"=","operand":"z39UzmN7GhPqkYiQfVWY4pYADMQzv7qWC"}]]}'
    
    
    
#### transferFrom
    
+ NEW VERSION
    
        ./marineV2 808 localhost:50061  from to amount label
        
        ./marineV2 808 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW z4MTiNvHs76DpwSH3gpivBuvuDynuhefy 8.86 etc
            (priv -> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)
        
        ./marineV2 808 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW z4MTiNvHs76DpwSH3gpivBuvuDynuhefy 99999999 btc --> Not enough balance
            (priv -> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)
        
        ./marineV2 808 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW z4MTiNvHs76DpwSH3gpivBuvuDynuhefy 3.2 btc
            (priv -> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)
            check : ./marineV2 888 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW
            check : ./marineV2 888 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy
        
        (ED -> ECC)
        ./marineV2 808 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW z2UPj4Vr2pvREww7jjX1SDoqLtix7toVN 18 btc
        (priv -> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)
        
        (ED -> RSA)
        ./marineV2 808 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW z3W7tjpeVMkgr8EjVbCjsUcYNSh6t8fxs 8.08 btc
        (priv -> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)
        
        (ECC -> ED/RSA)
        ./marineV2 808 localhost:50061 z2UPj4Vr2pvREww7jjX1SDoqLtix7toVN z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW 1.01 btc
        ./marineV2 808 localhost:50061 z2UPj4Vr2pvREww7jjX1SDoqLtix7toVN z3W7tjpeVMkgr8EjVbCjsUcYNSh6t8fxs 2 btc
        (priv -> vDJgRQ3rKt2uZCSmR5QdKt6eGp7zFT2zDTAJTs52AMdjLiuowQtw3ea2wT87Cf8Lwqe2WUAoc8png3mp764f4HNXRdH6n3T5Zo1tf7AAaoBf65Fhek4vVmDGjxUHfc9ksXvgXRahruYYFGxU4Er1pYWnnyU5dZmjwZxwr)
        
        (RSA -> ED)
        ./marineV2 808 localhost:50061 z3W7tjpeVMkgr8EjVbCjsUcYNSh6t8fxs z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW 8.08 btc
        (priv -> 2rm2QNNaHzhtdRY7bCNXK9Wu7P8JUx5WrpUq6yAh7J6Aonukp4NgNEyE3p4w6xyvErPrVoPWA8EryazsGqUbAxEY9k8TuEro2J5MedrBdbq17AapWMKpZwsmes1EDuq1iDSGtjftfpAr6gLFyVJDfzo2WeHpUshxW467D2vHxQyqdTmkkKPEYRjh1TfoZxUU1e7xCinoNEfWfduz11E2BciSN2sX9h9YEhWRQj4nPwRs9AsoHykAA2USM3ru3fv1DCTGea5mVLyr6kTjqBrGSiCARmdBczGQD7P2ngstcyLKmSEA6mh8xF63Xwb5aeZRTmWJrqKc1yGXtM9uwuZjoT3raH7B1BP12ygnGw2rABdnWBZaNWuWxaKoTquMiQeq5jZRDduUGypG4aWNwL5RCWFkJEDAs7gdvTtWZDwxvbY4shjxWLCpQrPVAtG3zs7awWrP92tBGhJqfft7rojahfqhjP8UkkKSqJCCD8j5LyDgLMxRbiScUuorGbWCSSKo9XXZW7qCB5DwyV73v2qpMLgMxrLrFBrEzCdz585rBupJza2pZk4wab7xqEVwmKqj76UyDca6TaE89XSHL2XBBXKiPmTVneorGc746pgC12LXFkw8fFE2Gas6Nk4LQaj662iqSwRh7z1QSN4t6QdgMVxP2JfDfRoBHJJCSwHNUVQyCihjWmzyAQjm12yZLLwFxf8ZbSEzwyZH62Umb1XgQMh8Nzu2hXjMHfnjp852gbCf3uR3MBhGFNSG3rq5L5X46oUDUYZCkCR85ySXB3pPVmDaYUpunvD7QRbc1c5yuZZdb4fbYfbrKLENrMSp6GHYg3UCZk3R9uPFPtugNDZmovf4Fu2JV9WH6Mn8x6paCryXzk4W9gBj1rgvq3CSxqKdPHJUdkNeaKaQBxUd3U9HZrSXvXoH7CPFDZwceSh5FKaqqMtpucXRvDM3VxV6cgJFFTjt8GEzQeBaouexoMZdV52Fkermx6wQpU5RTckFxTo5zEjqLQiWFRtHWi1zUUDxp1XdGAFZELKju1eBKWrdWo3g1G58DCouxGpRkbXmVn2nK1X4jjbHDVk1woQY2EYmJmt8D6kwBhWTbgnVpqDNdDCYsZ4cUv6bPXWGtym4YHXxS9Ekdww3vMkhfZPcjVDtDbeh6UcFLQbSkKFGRXk3io2fg5iRLDWGDGAVtZM6HvSc6MUvf8F3Y3e4Diiuhrp9BZCcTfXGdX4ssR9y3sBxBAnYmdWemiLvrUauxRUnbnR7bHocs2N4EfJkrv794HRp5NjuCL3AGt8RT8jah3EDTTEf2oUxqVhGSP31nfV2RWwmdytp2XvGHvD2r5xGtA3wkG2k8xzebgvnxJC7amUXQfNXSwmidaJd2updYmZHxkdjaday9RJcd9GXoYozfu5c15uoH4U8U2ybqUc83WQdd7JfWzbmR6vDgRdhg9iqSREgugtmEhYmSPqMAAC2E6Jd4Se7buqhaMNdmRvUKoiHNvb6kNTz7MMbMrZhBGG8SFTYSkBewcLTR9NhCFTsddNt3Wn7P2NFS4PeUtix4VvbZEhPE1YzagpoUpzbvUAsB6G7NKqvCX2aYErjTfAH6RMQtdVHe55xoB6t9EFY9wCW6YoApxF)
        
        ./marineV2 808 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy z5cXux681QCFJqYdTaDju7Hp4NFQ7jCZo 70 etc
            (5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP)
        
        ./marineV2 808 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW 9 btc
            (priv -> 5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP)
        
            (transfer)
        ./marineV2 808 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3 6 btc
    
            (transfer)
        ./marineV2 808 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3 66 btc
    
            (balanceof)
        ./marineV2 888 localhost:50061 zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3

        ./marineV2 808 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW z5cXux681QCFJqYdTaDju7Hp4NFQ7jCZo 1000 btc
        
        ./marineV2 808 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW zkRuxEW3dEAyGBZaSEbdc8ZqseUUjvxo 888 btc
    
    
+ New tranferFrom2()

        ./marineV2 808 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW zkRuxEW3dEAyGBZaSEbdc8ZqseUUjvxo 80 btc
        (priv -> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)
    

+ fransferfrom() - Log Information

        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"to","operator":"=","operand":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW"}]]}'
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"from","operator":"=","operand":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW"}]]}'
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"to","operator":"=","operand":"z4MTiNvHs76DpwSH3gpivBuvuDynuhefy"}]]}'

#### Approve

    SCEBNARIO : *Jay* APPROVE *Paul* to move assets
    
        Jay : z4MTiNvHs76DpwSH3gpivBuvuDynuhefy (owner)
    
        Paul : z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW (spender1, approved now)
        
        Tim : zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3 (spender2, not approved yet --> plz check)

+ Admin Version with invoke

        ./marineV2 8 localhost:50061 Admin starpoly cat 
                '{"Args":["approve",{"owner":"address1", "spender":"address2", "lable":"btc", "limit" :"1000000", "expired":""}]}' pad

        ./marineV2 8 localhost:50061 Admin starpoly cat 
                '{"Args":["approve",{"owner":"z12abcdef", "spender":"zlkjhgffd", "lable":"btc", "limit" :"1000000", "expired":""}]}' 


+ Checks : Admin Version with Query

        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"spender","operator":"=","operand":"zlkjhgffd"}]]}'
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"spender","operator":"=","operand":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW"}]]}'
                
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"owner","operator":"=","operand":"z4MTiNvHs76DpwSH3gpivBuvuDynuhefy"}]]}'
    

+ Normal Version with signature check by invoke (818 --> 8 in marineDokcing)

        Command ip:port 818 owner spender label limit expired
            
        ./marineV2 818 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW btc 500 ""
        
        
        ./marineV2 818 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW "" 500 ""
        
        
        ./marineV2 818 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3 "" 500 ""
        (priv -> 5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP)


    
+ Approve()/Update()

        ./marineV2 818 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW btc 100 "expired"
            (privKey : 5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP : Jay Key)

+ allowed TransferFrom()

        ./marineV2 808 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW 1.88 btc
            (privKey : enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU : Paul Key)
            
        ./marineV2 808 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW 2 btc
            (privKey : 5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP : Tim Key)
    
    
    
#### **NOT Allowed** Transfer
    
+ Tim is **NOT** Allowed for others

        ./marineV2 808 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW 9 btc
            (privKey : 38UJ99hwPzmLZbxhmcXhuDoWZEbaGF4g6U78TuYdHjYi1pCZehpvVPuCrW7G9zh2cMzQpmK7z5xrPgCaQpkMUiMa : Tim Key)
        
+ Not allowed for **"SPX"**
    
        ./marineV2 808 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW 2 spx
            (privKey : enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU : Paul Key)



+ Check Approve TX
    ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"spender","operator":"=","operand":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW"}]]}'

    
    
    
+ Scenario
    +   Transferring from not-owned account is allowed with "approve()"
    +  Transdfer money from:[**z4MTiNvHs76DpwSH3gpivBuvuDynuhefy**] --> to:[**z4MTiNvHs76DpwSH3gpivBuvuDynuhefy**]}
        
            ./marineV2 808 localhost:50061  fromAddress toAddress amount label
            ./marineV2 808 localhost:50061  z4MTiNvHs76DpwSH3gpivBuvuDynuhefy z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW 9 btc
            (privKey : enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)


#### ALLOWANCE

+   Query allowance status with (onwer/spender)
    
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"spender","operator":"=","operand":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW"}]]}'

        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"owner","operator":"=","operand":"z4MTiNvHs76DpwSH3gpivBuvuDynuhefy"}]]}'

        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"spender","operator":"=","operand":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW"},{"field":"owner","operator":"=","operand":"z4MTiNvHs76DpwSH3gpivBuvuDynuhefy"}]]}'

+   Query allowance status with (onwer/spender) when you know both owner, spender addresses and label
    +  This function is used in the function of transferFrom()
    
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","z4MTiNvHs76DpwSH3gpivBuvuDynuhefy.z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW.btc"]}'
    
    + for allow to move all types of label in the account address
        ??? 
    
    
#### ISSUE

+ ISSUE_ only for Admin (undergoing,.. Not working)

    ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["issue",{"symbol":"btc", "issuer":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW", "totalsupply":"1000000000.0"}]}'
    
    ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["issue",{"symbol":"spx", "issuer":"z4MTiNvHs76DpwSH3gpivBuvuDynuhefy", "totalsupply":"1000000000"}]}'

+ ISSUE2 (working, but still under testing)

        ./marineV2 828 localhost:50061 symbol issuer totalSupply
        
        ./marineV2 828 localhost:50061 btc z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW 100000000
            (pk --> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)
        
        ./marineV2 828 localhost:50061 etc z4MTiNvHs76DpwSH3gpivBuvuDynuhefy 100000000
            (pk --> 5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP)
    
        ./marineV2 828 localhost:50061 xrp z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW 100000000

            (pk --> enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)

    - checks on ISSUE2 (NEW VERSION)
        
        - Check for "issue"  --> "WHOS ISSUER"
            (check for "establish")

        - Simple Checks (can be only by the issuer or someone in the membership)
    
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","z4MTiNvHs76DpwSH3gpivBuvuDynuhefy@etc"]}'
            
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"issuer","operator":"=","operand":"z4MTiNvHs76DpwSH3gpivBuvuDynuhefy"}]]}'
            
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"issuer","operator":"=","operand":"z4MTiNvHs76DpwSH3gpivBuvuDynuhefy"}]]}'
            
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"symbol","operator":"=","operand":"etc"}]]}'
            
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"founder","operator":"=","operand":"z4MTiNvHs76DpwSH3gpivBuvuDynuhefy"}]]}'
                (actually, this is derived from establish()..)
            
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","etc"]}'
                (actually, this is derived from establish()..)
        
        - Checks my status()
            

#### Whos Issuer
    
    ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"symbol","operator":"=","operand":"btc"}]]}'
    
    ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"symbol","operator":"=","operand":"spx"}]]}'
    
    ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"symbol","operator":"=","operand":"etc"}]]}'
    
    ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"issuer","operator":"=","operand":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW"}]]}'
    ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","z4MTiNvHs76DpwSH3gpivBuvuDynuhefy@etc"]}'
    
    ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"issuer","operator":"=","operand":"z4MTiNvHs76DpwSH3gpivBuvuDynuhefy"}, {"field":"symbol","operator":"=","operand":"etc"}]]}'
    
    
+ Advnaced
    
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector": {"symbol": "etc"},"fields": ["issuer","totalsupply"]}]}' --> working
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector": {"symbol": "etc"},"fields": ["issuer"]}]}' --> working
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector": {"symbol": "btc"},"fields": ["issuer"]}]}' --> working
    
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector": {"symbol": "etc"},"fields": ["mintedamount"]}]}' --> working
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"fields": ["totalsupply", "issuer"]}]}' 
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector": {"totalsupply":{"$gt":null}}, "fields": ["totalsupply", "issuer"]}]}'
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector": {"issuer":{"$gt":null}}}]}'
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector": {"issuer":{"$gt":null}}}]}'
        
        
    
    --> working but looks wierd

        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector":{"symbol":"etc"}}]}'  --> working




#### SPCIAL QURERY for grepping any values in the choosen fileds with the given conditions

        (all issuers)
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector": {"issuer":{"$gt":null}}}]}'
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector": {"symbol":{"$gt":null}}}]}'
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector": "$and":[{"mintedamount":{"$gt":null}}, {"minter":"zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3"}]}]}'  

        (Grep all values of any mintedamount with that minter is "zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3")
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector":{"$and":[{"mintedamount":{"$gt":null}}, {"minter":"zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3"}]}}]}'


#### MINT (undergoing,.. --> DONE) - keyword : "minter"

+   Usage

    ./marineV2 868 localhost:50061 minter symbol mintedamount
    
    ./marineV2 868 localhost:50061 zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3 etc 1000
    
        (priv : 38UJ99hwPzmLZbxhmcXhuDoWZEbaGF4g6U78TuYdHjYi1pCZehpvVPuCrW7G9zh2cMzQpmK7z5xrPgCaQpkMUiMa)

        JSON --> '{"Args":["mint",{"address":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW", "symbol":"etc", "mintedamount":"10000"}]}'
    

**Scenario**

+ Please Try with a new account not in the membership for tests as followings,

        issuer of "etc"     : z4MTiNvHs76DpwSH3gpivBuvuDynuhefy
        minter (Tim)        : zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3 
        
+   1) Minter(Tim) try to mint 1000 etc --> MUST FAIL cause Tim is not in the membership

        ./marineV2 868 localhost:50061 zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3 etc 1000
        (priv : 38UJ99hwPzmLZbxhmcXhuDoWZEbaGF4g6U78TuYdHjYi1pCZehpvVPuCrW7G9zh2cMzQpmK7z5xrPgCaQpkMUiMa)

+   2) Join [Tim] into [etc] by issuer(z4MTiNvHs76DpwSH3gpivBuvuDynuhefy)

        ./marineV2 102 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3 etc 2
            --> (private key : 5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP)
            --> result :    (key) zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3@etc
            
    - join check : 
            
            ./marineV2 105 localhost:50061 zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3 etc
            --> (parent private key : 5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP)
            --> (my private key : 38UJ99hwPzmLZbxhmcXhuDoWZEbaGF4g6U78TuYdHjYi1pCZehpvVPuCrW7G9zh2cMzQpmK7z5xrPgCaQpkMUiMa)

+ 3) Try again to mint 909(amount) in etc by minter (memebr but not the issuer)

        ./marineV2 868 localhost:50061 zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3 etc 909
        (priv : 38UJ99hwPzmLZbxhmcXhuDoWZEbaGF4g6U78TuYdHjYi1pCZehpvVPuCrW7G9zh2cMzQpmK7z5xrPgCaQpkMUiMa)
    
    - check mint!!

            (this is for mint information, log)
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"minter","operator":"=","operand":"zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3"}]]}' 
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"minter","operator":"=","operand":"zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3"}]]}' 
            
            (now, check real balance)
            ./marineV2 888 localhost:50061 zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3
            
            (check the updates of total supply)
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"symbol","operator":"=","operand":"etc"}]]}'
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"symbol","operator":"=","operand":"etc"}]]}'
            
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"issuer","operator":"=","operand":"z4MTiNvHs76DpwSH3gpivBuvuDynuhefy"}]]}'

+ 4) Try again to mint 777(amount) in etc by Non-member

        ./marineV2 868 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW etc 777
        (priv : enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)

+ 5) Try again to mint 777(amount) in etc by issuer, the money is minted at issuer(minter=issuer)
        
        ./marineV2 868 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy etc 777
        (priv : 5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP)
        
        (now, check real balance)
        ./marineV2 888 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy
        
+ 6) Minter Address Must be the same to the address that is used in signing
        
        NOTICE : The minter address should be issuer or member of the title(symbol)
    
+ mint check
            
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"minter","operator":"=","operand":"zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3"}]]}'   
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3@etc"]}'

+ [IMPORTANT] Update check for total supply --> undergoing!!
        
        IMPORTANT!!!
        The minted amount MUST be updated in the initial totalsupply 
        
    
#### BURN    - keyword : "burner"

+   Usage

        ./marineV2 848 localhost:50061 burner symbol burnedamount
        ./marineV2 848 localhost:50061 Admin starpoly cat '{"Args":["burn",{"address":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW", "symbol":"btc", "burnedamount":"10000"}]}'
        
        ./marineV2 848 localhost:50061 zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3 etc 888
        (current issued+minted : 100001908.00 - 888 (burned) = MUST BE 100,001,020.0000)
            (SnJT3 priv : 38UJ99hwPzmLZbxhmcXhuDoWZEbaGF4g6U78TuYdHjYi1pCZehpvVPuCrW7G9zh2cMzQpmK7z5xrPgCaQpkMUiMa) - member
            
        ./marineV2 848 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW etc 880.8  --> fail cuase non-member try!!
            (qdDXW priv : enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU) - non-member
        
            
        ./marineV2 848 localhost:50061 zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3 etc 880.8  -> fail cuase burner is not signer
            (qdDXW priv : enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU) -> burner != signer  
            
+ checks
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector":{"$and":[{"burnedamount":{"$gt":null}}, {"burner":"zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3"}]}}]}'
        
        (Grep all of any burn)
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector":{"$and":[{"burnedamount":{"$gt":null}}]}}]}'
        
        
        (Grep all if you have "title")
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector":{"title":{"$gt":null}}}]}'

        (Grep all if you have "title of etc")       ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector":{"$and":[{"title":{"$gt":null}}, {"title":"etc"}]}}]}'
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector":{"$and":[{"title":"etc"}]}}]}'
        
    

#### Lookup Log
    
    ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"from","operator":"=","operand":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW"}]]}'
    ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"from","operator":"=","operand":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW"}]]}'
    
    ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"to","operator":"=","operand":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW"}]]}'
    
    
    ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"from","operator":"=","operand":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW"}]]}'
    

#### Check Log
    ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"issuer","operator":"=","operand":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW"}]]}'
    
    ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"issuer","operator":"=","operand":null}]]}'


#### DoubleChecks
    ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"from","operator":"=","operand":"z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW"}]]}'
    afe8f81c75f5733e6bc184dce3e93e5229df5954d7ace9e11f2634408a6795a9



## EXCHANGE

#### eXchange()

    ./marineV2 200 localhost:50061 exchanger, from, to, fromlabel, tolabel, excamount)



##Authentication Control Engine (ACE)

+ **ESTABLISH**

    ./marineV2 101 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW btc
    
        --> (private key : enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)

    - Check by Admin (if STATUS() has not made yet)
    
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","btc"]}'
        
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW@btc"]}'
    
        ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query",[{"field":"title","operator":"=","operand":"btc"}]]}'


+ **JOIN**

        ./marineV2 102 localhost:50061 host guest title type {public(0), private(1), memberOnly(2), left(4)}
    
        (TEST  --> MUST BE success!!)
        ./marineV2 102 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW z4MTiNvHs76DpwSH3gpivBuvuDynuhefy btc 0
            --> (private key(DXW) : enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)
            
        ./marineV2 102 localhost:50061 z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW zXQEzunzC3SxUDt3AdyUuhXAPjDSnJT3 btc 0
            --> (private key(DXW) : enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)
        
        (TEST  --> MUST BE FAILED!!)
        ./marineV2 102 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy z4MTiNvHs76DpwSH3gpivBuvuDynuhefy btc 0
            --> (private key : 5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP)


    - check!!
            
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","z2Y6dzY3igFt5J5V2u2NDcSNXEFYqdDXW@btc"]}'
            
                - if want to know the founder's info, --> ask with titleInfo --> '{"Args":["queryByKey","btc"]}' --> u can find the whole info
            
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","z4MTiNvHs76DpwSH3gpivBuvuDynuhefy@btc"]}'
            
            (del to test again)
            ./marineV2 8 localhost:50061 Admin starpoly cat '{"Args":["delete","z4MTiNvHs76DpwSH3gpivBuvuDynuhefy@btc"]}' opt
            
+ **LEAVE**
    
        ./marineV2 103 localhost:50061 member title
    
    + Leave by himself
        
            ./marineV2 103 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy btc
            >> PK : 5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP)
        
    + Other member(paul) kicks him out form the membership
        
            ./marineV2 103 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy btc
        
            >> PK : enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU)

        - check
             ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","z4MTiNvHs76DpwSH3gpivBuvuDynuhefy@btc"]}'

+ **MODIFYMEMBER**

    + ONLY ALLOWED TO MODIFY [TYPE : public(0), private(1), memberOnly(2), left(4, reserved)]
    
            ./marineV2 104 localhost:50061 member title type
        
            ./marineV2 104 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy btc 0
        
                (PK : enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU --> parent)
                (PK : 5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP --> me)

+ **STATUS**

    - Admin check

            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByKey","z4MTiNvHs76DpwSH3gpivBuvuDynuhefy@btc"]}'
            
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["query","z4MTiNvHs76DpwSH3gpivBuvuDynuhefy@btc"]}'

    - Only parent(host) and I(guest) can look up the branch

            ./marineV2 105 localhost:50061 memberAddress
            
            ./marineV2 105 localhost:50061 z4MTiNvHs76DpwSH3gpivBuvuDynuhefy btc
            
            (PK : 5db9G8PZMFT65JbsGhwdf5UtVGgdmemiTXxwWu8dpUrJHrRiiGCtCgNiTMAEJRHZwus9869UJLpwXqB2ymZzdFZP 
                            --> me --> MUST WORK!!)
            (PK : enxYgKm2Cf71EsK7osNrujuUqoUjuwEFy1GAwyVarvnuDhme3hDjuvXYSmbyFPaLUfyeZdHDHefwgb75ET9HmoU 
                            --> parent(host) --> MUST WORK!!)
            (PK : 38UJ99hwPzmLZbxhmcXhuDoWZEbaGF4g6U78TuYdHjYi1pCZehpvVPuCrW7G9zh2cMzQpmK7z5xrPgCaQpkMUiMa 
                            --> someone --> MUST FAIL(if not public!!), Success(if public)

    - How to grep all values in the given "title"??

            if you passed the basic status() above with a given title, now you are allowed to call Admin to grep all information, but it would be only onetime.
            
            --> Need to implement to grep all values in the given title if in the member?
            --> ./marineV2 106 localhost:50061 all title (expected)
            
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector":{"$and":[{"title":"etc"}]}}]}'
            
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector":{"title":"etc"}}]}'
            
            ./marineV2 7 localhost:50061 Admin starpoly cat '{"Args":["queryByPass",{"selector":{"title":{"$gt":null}}}]}'
            

    
    