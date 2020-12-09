package eu.fasten.analyzer.complianceanalyzer;

import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.io.InputStream;

import org.apache.commons.io.IOUtils;
import org.json.simple.JSONArray;
import org.json.simple.JSONObject;
import org.json.simple.parser.JSONParser;
import org.json.simple.parser.ParseException;




public class Main {

    public static void main(String args[]) throws IOException, ParseException {

        // Reading the input JSON file in the `resources` directory
    	
       try (InputStream is = Thread.currentThread().getContextClassLoader().getResourceAsStream("inputKafkaTopic.json")) {

		    	String reader = IOUtils.toString(is, "UTF-8");   			
		       	
		       	JSONParser jsonparser=new JSONParser();
		   		Object obj=jsonparser.parse(reader);
		   		JSONObject empjsonobj=(JSONObject)obj;
		       	
		   		String pname=(String) empjsonobj.get("plugin_name");
		   		String pver=(String) empjsonobj.get("plugin_version");
		   		
		   		System.out.println("Plugin Name:"+pname);
		   		System.out.println("Plugin Version:"+pver+"\n");
		   		
		   		JSONObject inputobj=(JSONObject)empjsonobj.get("input");
		   		
		   		String product=(String) inputobj.get("product");
		   		String forge=(String) inputobj.get("forge");
		   		String version=(String) inputobj.get("version");
		   		
		   		System.out.println("Input field details are :\n");
		   		
		   		System.out.println("Product:"+product);
		   		System.out.println("Forge:"+forge);
		   		System.out.println("Version:"+version+"\n");
		   		
		   		
		   		Number create=(Number) empjsonobj.get("created_at");
		   		System.out.println("Created at:"+create+"\n");
		   		
		   		
		   		JSONObject payload=(JSONObject)empjsonobj.get("payload");
		   		JSONObject pmeta=(JSONObject)payload.get("package_metadata");
		   		JSONArray larray=(JSONArray)pmeta.get("licenses");	
		   		JSONArray auarray=(JSONArray)pmeta.get("authors");	
		   		System.out.println("Package_metadata details are :\n");
		   		for (int i=0;i<larray.size();i++)
		   		{
		   			String lic=(String) larray.get(i);
		   			String auth=(String) auarray.get(i);
		   			System.out.println("---"+"Auther "+auth+" License is :"+lic+"\n");
		   		}
		   		
		   		System.out.println("File_metadata details are :\n");
		   		
		   		JSONArray fmeta=(JSONArray)payload.get("file_metadata");
		   		for (int i=0;i<fmeta.size();i++)
		   		{
		   			
		   			JSONObject fileobj=(JSONObject) fmeta.get(i);
		   			String file_name=(String) fileobj.get("filename");
		   			String sha=(String) fileobj.get("SHA-1");
		   			System.out.println("---FileName: "+file_name);
		   			System.out.println("---SHA-1: "+sha);
		   			JSONArray lice=(JSONArray) fileobj.get("licenses");
		   			
		   			for(int j=0;j<lice.size();j++)
		   			{
		   				  String license=(String)lice.get(j);
		   				  System.out.println("---License: "+license+"\n");
		   			}
		   			
		   		}
		
		        }
     
    	
        catch (Exception e) {
            throw new RuntimeException(e);
        } 
    }
}
