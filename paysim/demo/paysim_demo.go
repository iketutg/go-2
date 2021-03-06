package demo

var Demo_Specs string =`


{ "Specs":[{"SpecName": "ISO8583_1_v1__DEMO_", "Fields": [{"BitPosition":0,"Name": "Message Type", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":4,"DataEncoding":"ascii","Padding":"default"}},
                                                           {"BitPosition":0,"Name": "Bitmap", "Type":"Bitmapped", "Attrs":{"DataEncoding":"binary"},"Children":
                                                          [{"BitPosition":2,"Name": "PAN", "Type":"Variable", "Attrs":{"Key":false,"FieldIndicatorLength":2,"FieldIndicatorEncoding":"ascii","DataEncoding":"ascii","Padding":"default"}},
                                                           {"BitPosition":3,"Name": "Processing Code", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":6,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":4,"Name": "Amount, Transaction", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":12,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":5,"Name": "Amount, Settlement", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":12,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":6,"Name": "Amount, Cardholder Billing", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":12,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":7,"Name": "Date and Time, Transmission", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":10,"DataEncoding":"ascii","Padding":"default"}},
                                                           
                                                           {"BitPosition":8,"Name": "Amount, cardholder billing fee", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":8,"DataEncoding":"ascii","Padding":"default"}},
                                                           {"BitPosition":9,"Name": "Conversion rate, settlement", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":8,"DataEncoding":"ascii","Padding":"default"}},
                                                           {"BitPosition":10,"Name": "Conversion rate, cardholder billing", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":8,"DataEncoding":"ascii","Padding":"default"}},
                                                           
                                                           {"BitPosition":11,"Name": "STAN", "Type":"Fixed", "Attrs":{"Key":true,"FieldLength":6,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":12,"Name": "Time, local transaction (hhmmss)", "Type":"Fixed", "Attrs":{"Key":true,"FieldLength":6,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":13,"Name": "Date, local transaction (MMDD)", "Type":"Fixed", "Attrs":{"Key":true,"FieldLength":4,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           
                                                           {"BitPosition":14,"Name": "Date, Expiry", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":4,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           
                                                           {"BitPosition":15,"Name": "Date, Settlement", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":4,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":16,"Name": "Date, Conversion", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":4,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":17,"Name": "Date, Capture", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":4,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           
                                                           {"BitPosition":18,"Name": "Merchant, Type", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":4,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":19,"Name": "Acquiring institution country code", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":3,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":22,"Name": "POS Entry Mode", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":3,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":23,"Name": "Application PAN Sequence Number", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":3,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":24,"Name": "Function Code", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":3,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":25,"Name": "Point of service condition code", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":2,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":26,"Name": "Point of service Capture code", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":2,"DataEncoding":"ascii","Padding":"default"}}, 
                                                           {"BitPosition":27,"Name": "Authorizing identification response length", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":1,"DataEncoding":"ascii","Padding":"default"}}, 
                                      
                                                           {"BitPosition":35,"Name": "Track 2", "Type":"Variable", "Attrs":{"Key":false,"FieldIndicatorLength":2,"FieldIndicatorEncoding":"ascii","DataEncoding":"ascii","Padding":"default"}},
                                                           {"BitPosition":38,"Name": "Approval Code", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":6,"DataEncoding":"ascii","Padding":"default"}},
                                                           {"BitPosition":39,"Name": "Action Code", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":3,"DataEncoding":"ascii","Padding":"default"}},
                                                           {"BitPosition":52,"Name": "PIN Data", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":8,"DataEncoding":"binary","Padding":"default"}},
                                                           {"BitPosition":55,"Name": "ICC Data", "Type":"Variable", "Attrs":{"Key":false,"FieldIndicatorLength":3,"FieldIndicatorEncoding":"ascii","DataEncoding":"binary","Padding":"default"}},
                                                           {"BitPosition":64,"Name": "MAC(1)", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":8,"DataEncoding":"binary","Padding":"default"}},
                                                           {"BitPosition":96,"Name": "Key Management Data", "Type":"Variable", "Attrs":{"Key":false,"FieldIndicatorLength":2,"FieldIndicatorEncoding":"ascii","DataEncoding":"binary","Padding":"default"}},
                                                           {"BitPosition":128,"Name": "MAC(2)", "Type":"Fixed", "Attrs":{"Key":false,"FieldLength":8,"DataEncoding":"binary","Padding":"default"}}
                                                          ]
                                                        }  
                                             ]
          }
          ]
}            `;
