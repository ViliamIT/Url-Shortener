# Url-Shortener
Server sa po spusteni pripaja na lokalnu databazu, stranka je na adrese http://localhost:8080/homepage po zadani url do fomulara server odosle naspat funkcnu skratenu url. Po prijati url server vytvori zaznam v databaze a vrati skratenu url napriklad http://localhost:8080/3 kde /3 je vlastne zaroven ID tej url adresy v databaze. Skratena url adresa po zadani do browsera nas presunie na originalnu stranku tak, ze podla ID co je za "/" vyhlada v databaze entitu ktora sa s nim zhoduje a spravi redirect na orgininalnu URL
