import{readFileSync as P}from"fs"
let l=P("i").toString().split("\r\n")
let r=l.splice(0,1)[0].split(",")
let s=[]
let p=[]
for(let q of l){if(!q){if(p.length!=0){s.push(p)
p=[]}continue}p.push(q.replaceAll("  "," ").trim().split(" "))}
for(let a of r){s=s.map(u=>u.map(o=>o.map(t=>t==a?"X":t)))
let f=s.find((y,d)=>y.some(row=>5==row.reduce((v,c)=>c=="X"?v+1:v,0))||y.some((_,i)=>5==y.reduce((p,c)=>c[i]=="X"?p+1:p,0)))
if(f){console.log(parseInt(a)*f.reduce((e,c)=>e+c.reduce((e,c)=>e+(isNaN(parseInt(c))?0:parseInt(c)),0),0))
process.exit()}}
