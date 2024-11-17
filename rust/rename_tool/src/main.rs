use std::{env, fs};
use std::path::Path;
use regex::Regex;

fn main() -> std::io::Result<()> {
    let mut dir = env::args().nth(1).unwrap_or(String::from(""));
    if dir.len() == 0 {
        dir = String::from(env::current_dir()?.to_str().unwrap());
    }
    println!("dir is {}", dir);

    let re = Regex::new(r".+_(bin.+\.dat)").unwrap();

    let dest = Path::new(dir.as_str()).join("dest");
    let mut dest_ready = false;

    for entry in fs::read_dir(dir)? {
        let entry = entry?;
        let path = entry.path();
        if path.is_file() {
            let file_name = path.file_name().and_then(|name| name.to_str()).unwrap();
            // println!("Found: {}", file_name);
            let Some(caps) = re.captures(file_name) else {
                continue;
            };
            // println!("caps: {:?}", caps);
            if caps.len() == 2 {
                let new_file_name = caps.get(1).unwrap().as_str();
                println!("new_file_name: {:?}", new_file_name);
                if !dest_ready && !dest.exists() {
                    fs::create_dir_all(dest.clone())?;
                    dest_ready = true;
                }
                fs::rename(path.clone(), dest.join(new_file_name))?
            }
        }
    }
    Ok(())
}
