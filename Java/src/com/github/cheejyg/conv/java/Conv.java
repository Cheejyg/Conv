package com.github.cheejyg.conv.java;

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import com.github.cheejyg.conv.java.*;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Conversion.
 * 
 * @author Cheejyg
 */
public class Conv {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ToString
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ToString parses <code>Object</code> object_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param object_ The <code>Object</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static java.lang.String ToString(java.lang.Object object_) {
		return object_.toString();
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Array conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * BooleansToBooleanArray parses <code>Boolean[]</code> booleans_ to
	 * <code>Boolean[]</code>.
	 * 
	 * @author Cheejyg
	 * @param booleans_ The <code>Boolean[]</code> to be parsed.
	 * @return <code>Boolean[]</code>
	 */
	public static java.lang.Boolean[] BooleansToBooleanArray(java.lang.Boolean... booleans_) {
		return booleans_;
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Long conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * LongToString parses <code>long</code> long_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static java.lang.String LongToString(long long_) {
		String string_ = ToString(long_);

		return string_;
	}

	/**
	 * LongToLong parses <code>long</code> long_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static java.lang.Long LongToLong(long long_) {
		return long_;
	}

	/**
	 * LongToInteger parses <code>long</code> long_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static java.lang.Integer LongToInteger(long long_) {
		java.lang.Integer integer_ = new java.lang.Integer((int) long_);

		return integer_;
	}

	/**
	 * LongToShort parses <code>long</code> long_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static java.lang.Short LongToShort(long long_) {
		Short short_ = new Short((short) long_);

		return short_;
	}

	/**
	 * LongToByte parses <code>long</code> long_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static java.lang.Byte LongToByte(long long_) {
		java.lang.Byte byte_ = new java.lang.Byte((byte) long_);

		return byte_;
	}

	/**
	 * LongToBoolean parses <code>long</code> long_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Boolean</code>
	 */
	public static java.lang.Boolean LongToBoolean(long long_) {
		if (long_ == 0) {
			return false;
		} else if (long_ == 1) {
			return true;
		} else {
		}

		return null;
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Int conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * IntToString parses <code>int</code> int_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>String</code>
	 * @see com.github.cheejyg.conv.java.Int
	 * @see com.github.cheejyg.conv.java.Int#IntToString(int int_)
	 */
	public static java.lang.String IntToString(int int_) {
		return Int.IntToString(int_);
	}

	/**
	 * IntToLong parses <code>int</code> int_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Long</code>
	 * @see com.github.cheejyg.conv.java.Int
	 * @see com.github.cheejyg.conv.java.Int#IntToLong(int int_)
	 */
	public static java.lang.Long IntToLong(int int_) {
		return Int.IntToLong(int_);
	}

	/**
	 * IntToInteger parses <code>int</code> int_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Integer</code>
	 * @see com.github.cheejyg.conv.java.Int
	 * @see com.github.cheejyg.conv.java.Int#IntToInteger(int int_)
	 */
	public static java.lang.Integer IntToInteger(int int_) {
		return Int.IntToInteger(int_);
	}

	/**
	 * IntToShort parses <code>int</code> int_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Short</code>
	 * @see com.github.cheejyg.conv.java.Int
	 * @see com.github.cheejyg.conv.java.Int#IntToShort(int int_)
	 */
	public static java.lang.Short IntToShort(int int_) {
		return Int.IntToShort(int_);
	}

	/**
	 * IntToByte parses <code>int</code> int_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Byte</code>
	 * @see com.github.cheejyg.conv.java.Int
	 * @see com.github.cheejyg.conv.java.Int#IntToByte(int int_)
	 */
	public static java.lang.Byte IntToByte(int int_) {
		return Int.IntToByte(int_);
	}

	/**
	 * IntToBoolean parses <code>int</code> int_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Boolean</code>
	 * @see com.github.cheejyg.conv.java.Int
	 * @see com.github.cheejyg.conv.java.Int#IntToBoolean(int int_)
	 */
	public static java.lang.Boolean IntToBoolean(int int_) {
		return Int.IntToBoolean(int_);
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Short conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ShortToString parses <code>short</code> short_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static java.lang.String ShortToString(short short_) {
		String string_ = ToString(short_);

		return string_;
	}

	/**
	 * ShortToLong parses <code>short</code> short_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static java.lang.Long ShortToLong(short short_) {
		Long long_ = new Long(short_);

		return long_;
	}

	/**
	 * ShortToInteger parses <code>short</code> short_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static java.lang.Integer ShortToInteger(short short_) {
		java.lang.Integer integer_ = new java.lang.Integer(short_);

		return integer_;
	}

	/**
	 * ShortToShort parses <code>short</code> short_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static java.lang.Short ShortToShort(short short_) {
		return short_;
	}

	/**
	 * ShortToByte parses <code>short</code> short_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static java.lang.Byte ShortToByte(short short_) {
		java.lang.Byte byte_ = new java.lang.Byte((byte) short_);

		return byte_;
	}

	/**
	 * ShortToBoolean parses <code>short</code> short_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Boolean</code>
	 */
	public static java.lang.Boolean ShortToBoolean(short short_) {
		if (short_ == 0) {
			return false;
		} else if (short_ == 1) {
			return true;
		} else {
		}

		return null;
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Byte conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ByteToString parses <code>byte</code> byte_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static java.lang.String ByteToString(byte byte_) {
		String string_ = ToString(byte_);

		return string_;
	}

	/**
	 * ByteToLong parses <code>byte</code> byte_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static java.lang.Long ByteToLong(byte byte_) {
		Long long_ = new Long(byte_);

		return long_;
	}

	/**
	 * ByteToInteger parses <code>byte</code> byte_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static java.lang.Integer ByteToInteger(byte byte_) {
		java.lang.Integer integer_ = new java.lang.Integer(byte_);

		return integer_;
	}

	/**
	 * ByteToShort parses <code>byte</code> byte_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static java.lang.Short ByteToShort(byte byte_) {
		Short short_ = new Short(byte_);

		return short_;
	}

	/**
	 * ByteToByte returns <code>Byte</code> byte_.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static java.lang.Byte ByteToByte(byte byte_) {
		return byte_;
	}

	/**
	 * ByteToBoolean parses <code>byte</code> byte_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Boolean</code>
	 */
	public static java.lang.Boolean ByteToBoolean(byte byte_) {
		if (byte_ == 0) {
			return false;
		} else if (byte_ == 1) {
			return true;
		} else {
		}

		return null;
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Boolean conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * BooleanToString parses <code>boolean</code> boolean_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>String</code>
	 * @see com.github.cheejyg.conv.java.Boolean
	 * @see com.github.cheejyg.conv.java.Boolean#BooleanToString(boolean boolean_)
	 */
	public static java.lang.String BooleanToString(boolean boolean_) {
		return Boolean.BooleanToString(boolean_);
	}

	/**
	 * BooleanToLong parses <code>boolean</code> boolean_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Long</code>
	 * @see com.github.cheejyg.conv.java.Boolean
	 * @see com.github.cheejyg.conv.java.Boolean#BooleanToLong(boolean boolean_)
	 */
	public static java.lang.Long BooleanToLong(boolean boolean_) {
		return Boolean.BooleanToLong(boolean_);
	}

	/**
	 * BooleanToInteger parses <code>boolean</code> boolean_ to
	 * <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Integer</code>
	 * @see com.github.cheejyg.conv.java.Boolean
	 * @see com.github.cheejyg.conv.java.Boolean#BooleanToInteger(boolean boolean_)
	 */
	public static java.lang.Integer BooleanToInteger(boolean boolean_) {
		return Boolean.BooleanToInteger(boolean_);
	}

	/**
	 * BooleanToShort parses <code>boolean</code> boolean_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Short</code>
	 * @see com.github.cheejyg.conv.java.Boolean
	 * @see com.github.cheejyg.conv.java.Boolean#BooleanToShort(boolean boolean_)
	 */
	public static java.lang.Short BooleanToShort(boolean boolean_) {
		return Boolean.BooleanToShort(boolean_);
	}

	/**
	 * BooleanToByte parses <code>boolean</code> boolean_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Byte</code>
	 * @see com.github.cheejyg.conv.java.Boolean
	 * @see com.github.cheejyg.conv.java.Boolean#BooleanToByte(boolean boolean_)
	 */
	public static java.lang.Byte BooleanToByte(boolean boolean_) {
		return Boolean.BooleanToByte(boolean_);
	}

	/**
	 * BooleanToBoolean returns <code>Boolean</code> boolean_.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be returned.
	 * @return <code>Boolean</code>
	 * @see com.github.cheejyg.conv.java.Boolean
	 * @see com.github.cheejyg.conv.java.Boolean#BooleanToBoolean(boolean boolean_)
	 */
	public static java.lang.Boolean BooleanToBoolean(boolean boolean_) {
		return Boolean.BooleanToBoolean(boolean_);
	}
}
